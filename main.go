// Package main implements a Model Context Protocol (MCP) server that provides
// text completion capabilities using local LLama.cpp models.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	mcpgolang "github.com/metoro-io/mcp-golang"
	mcphttp "github.com/metoro-io/mcp-golang/transport/http"
)

// Configuration constants define application defaults and limits
const (
	// ShutdownTimeout is the maximum time to wait for graceful shutdown
	ShutdownTimeout = 30 * time.Second
	// DefaultConfigFile is the default environment configuration file name
	DefaultConfigFile = "byte-vision-cfg.env"
)

// CompletionMetrics tracks performance and usage statistics for completion requests
type CompletionMetrics struct {
	RequestCount  int64         // Total number of completion requests received
	SuccessCount  int64         // Number of successful completions
	ErrorCount    int64         // Number of failed completions
	TimeoutCount  int64         // Number of requests that timed out
	TotalDuration time.Duration // Cumulative time spent on all requests
	AverageTokens float64       // Average number of tokens generated per request
}

// Global variables for application configuration and state management
var (
	llamaCliArgs LlamaCliArgs   // Configuration for LLama.cpp command-line arguments
	appArgs      DefaultAppArgs // General application configuration
	logger       *log.Logger    // Custom logger instance for structured logging
	logFile      *os.File       // Handle to the log file for cleanup
	shutdownOnce sync.Once      // Ensures cleanup only happens once during shutdown
)

// CompletionArguments defines the input structure for the MCP completion tool
type CompletionArguments struct {
	Prompt string `json:"prompt" description:"The prompt text to generate completion for"`
}

// setupLogging configures dual logging to both file and console with structured output.
// It creates the logs directory if it doesn't exist and sets up a multi-writer logger.
//
// Returns:
//   - error: Any error that occurred during log setup
func setupLogging() error {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(appArgs.AppLogPath, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Construct the full path to the log file
	logFilePath := filepath.Join(appArgs.AppLogPath, appArgs.AppLogFileName)

	// Open the log file with creation, write, and append permissions
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create multi-writer to output to both console and file simultaneously
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Create a custom logger with [APP] prefix and timestamp/file information
	logger = log.New(multiWriter, "[APP] ", log.LstdFlags|log.Lshortfile)

	// Replace the default logger to capture all log output
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger.Printf("Logging initialized - writing to %s", logFilePath)
	return nil
}

// cleanup handles graceful resource cleanup during application shutdown.
// It uses sync.Once to ensure cleanup only happens once, even if called multiple times.
func cleanup() {
	shutdownOnce.Do(func() {
		if logFile != nil {
			logger.Println("Closing log file...")
			if err := logFile.Close(); err != nil {
				log.Printf("Error closing log file: %v", err)
			}
		}
	})
}

// main is the application entry point that handles initialization, server startup,
// and graceful shutdown coordination.
func main() {
	// Initialize basic logging for startup messages before full logging setup
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Ensure cleanup happens regardless of how the application exits
	defer cleanup()

	// Load environment variables from configuration file
	if err := godotenv.Load(DefaultConfigFile); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Parse configuration from environment variables
	llamaCliArgs = ParseDefaultLlamaCliEnv()
	appArgs = ParseDefaultAppEnv()

	// Setup structured logging to file and console
	if err := setupLogging(); err != nil {
		log.Fatalf("Failed to setup logging: %v", err)
	}

	logger.Println("Application starting...")

	// Create context for coordinating graceful shutdown across goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling for graceful shutdown (Ctrl+C, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the MCP server in a separate goroutine
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- runServer(ctx)
	}()

	// Wait for either a shutdown signal or server error
	select {
	case <-quit:
		logger.Println("Received shutdown signal...")
	case err := <-serverErr:
		if err != nil && !errors.Is(err, context.Canceled) {
			logger.Printf("Server error: %v", err)
		}
	}

	// Initiate graceful shutdown by canceling the context
	cancel()

	// Give the server time to shut down gracefully before forcing termination
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer shutdownCancel()

	select {
	case <-shutdownCtx.Done():
		logger.Println("Forced shutdown after timeout")
	case <-serverErr:
		logger.Println("Server shutdown complete")
	}

	logger.Println("Application shutdown complete")
}

// runServer initializes and starts the MCP HTTP server with the completion tool.
// It handles server lifecycle management and graceful shutdown coordination.
//
// Parameters:
//   - ctx: Context for server lifecycle management and cancellation
//
// Returns:
//   - error: Any error that occurred during server operation
func runServer(ctx context.Context) error {
	// Create HTTP transport for the MCP protocol
	transport := mcphttp.NewHTTPTransport(appArgs.EndPoint)
	transport.WithAddr(appArgs.HttpPort)

	// Create the MCP server instance
	server := mcpgolang.NewServer(transport)

	// Register the text completion tool with the server
	if err := server.RegisterTool(
		"generate_completion",
		"Generate text completion using the local LLM",
		handleCompletionTool,
	); err != nil {
		return fmt.Errorf("failed to register completion tool: %w", err)
	}

	logger.Printf("Starting MCP HTTP server on %s%s", appArgs.HttpPort, appArgs.EndPoint)

	// Start the server in a separate goroutine to allow for cancellation
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Serve()
	}()

	// Wait for either context cancellation or server error
	select {
	case <-ctx.Done():
		logger.Println("Shutting down server...")
		// Attempt graceful transport shutdown
		if err := transport.Close(); err != nil {
			logger.Printf("Transport shutdown error: %v", err)
		}
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

// handleCompletionTool processes MCP completion requests by validating input,
// executing LLama.cpp, and returning formatted responses with error handling.
//
// Parameters:
//   - arguments: The completion request containing the prompt text
//
// Returns:
//   - *mcpgolang.ToolResponse: Formatted response containing the completion or error
//   - error: Any error that occurred during request processing
func handleCompletionTool(arguments CompletionArguments) (*mcpgolang.ToolResponse, error) {
	// Initialize metrics tracking for this request
	var metrics CompletionMetrics
	startTime := time.Now()
	metrics.RequestCount++

	// Track request duration and log performance metrics
	defer func() {
		duration := time.Since(startTime)
		metrics.TotalDuration += duration
		logger.Printf("Request completed in %v (avg: %v)",
			duration,
			time.Duration(int64(metrics.TotalDuration)/metrics.RequestCount))
	}()

	// Validate that the prompt is not empty
	if arguments.Prompt == "" {
		logger.Println("Empty prompt received")
		return &mcpgolang.ToolResponse{
			Content: []*mcpgolang.Content{
				mcpgolang.NewTextContent("Error: Prompt cannot be empty"),
			},
		}, nil
	}

	// Log the incoming request with truncated prompt for readability
	logger.Printf("Handling completion request for prompt: %.100s...", arguments.Prompt)

	// Get timeout configuration with fallback to default
	timeoutSeconds := appArgs.TimeOutSeconds
	if timeoutSeconds <= 0 {
		timeoutSeconds = 300 // fallback default of 5 minutes
	}

	// Create context with timeout for the completion request
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	logger.Printf("Starting completion with timeout of %d seconds", timeoutSeconds)

	// Prepare command-line arguments for LLama.cpp using configuration
	args := prepareLlamaArgs(arguments.Prompt)

	// Execute the completion generation
	output, err := GenerateSingleCompletionWithCancel(ctx, appArgs, args)
	if err != nil {
		// Handle timeout errors specifically
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Printf("Completion timed out after %d seconds", timeoutSeconds)
			return &mcpgolang.ToolResponse{
				Content: []*mcpgolang.Content{
					mcpgolang.NewTextContent(fmt.Sprintf("Error: Completion timed out after %d seconds", timeoutSeconds)),
				},
			}, nil
		}

		// Handle other execution errors
		logger.Printf("Error generating completion: %v", err)
		return &mcpgolang.ToolResponse{
			Content: []*mcpgolang.Content{
				mcpgolang.NewTextContent(fmt.Sprintf("Error generating completion: %v", err)),
			},
		}, nil
	}

	logger.Printf("Completion generated successfully, output length: %d chars", len(output))

	// Return successful completion as MCP tool response
	return &mcpgolang.ToolResponse{
		Content: []*mcpgolang.Content{
			mcpgolang.NewTextContent(string(output)),
		},
	}, nil
}

// prepareLlamaArgs constructs command-line arguments for LLama.cpp by combining
// configuration from environment variables with the user-provided prompt.
// It filters out any existing prompt arguments to avoid conflicts.
//
// Parameters:
//   - prompt: The user-provided prompt text to include in the arguments
//
// Returns:
//   - []string: Complete command-line arguments array for LLama.cpp
func prepareLlamaArgs(prompt string) []string {
	// Start with base arguments from environment configuration
	args := LlamaCliStructToArgs(llamaCliArgs)

	// Filter out any existing --prompt arguments to avoid conflicts
	filteredArgs := make([]string, 0, len(args))
	skipNext := false

	for _, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}
		if arg == "--prompt" {
			skipNext = true // Skip the next argument (the old prompt value)
			continue
		}
		filteredArgs = append(filteredArgs, arg)
	}

	// Add the new prompt as the final arguments
	filteredArgs = append(filteredArgs, "--prompt", prompt)

	logger.Printf("Prepared llama args with %d parameters", len(filteredArgs))
	return filteredArgs
}
