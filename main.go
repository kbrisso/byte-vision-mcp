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

// Configuration constants
const (
	ShutdownTimeout   = 30 * time.Second
	DefaultConfigFile = "byte-vision-cfg.env"
)

// Add metrics tracking
type CompletionMetrics struct {
	RequestCount  int64
	SuccessCount  int64
	ErrorCount    int64
	TimeoutCount  int64
	TotalDuration time.Duration
	AverageTokens float64
}

// Global variables for configuration
var (
	llamaCliArgs LlamaCliArgs
	appArgs      DefaultAppArgs
	logger       *log.Logger
	logFile      *os.File
	shutdownOnce sync.Once
)

// CompletionArguments defines the input structure for the completion tool
type CompletionArguments struct {
	Prompt string `json:"prompt" description:"The prompt text to generate completion for"`
}

// setupLogging configures file logging to the logs directory
func setupLogging() error {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(appArgs.AppLogPath, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	logFilePath := filepath.Join(appArgs.AppLogPath, appArgs.AppLogFileName)

	// Open log file
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create multi-writer to write to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Create custom logger
	logger = log.New(multiWriter, "[APP] ", log.LstdFlags|log.Lshortfile)

	// Replace default logger
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger.Printf("Logging initialized - writing to %s", logFilePath)
	return nil
}

// cleanup handles resource cleanup
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

func main() {
	// Initialize a basic logger first for startup errors
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Ensure cleanup happens
	defer cleanup()

	// Load environment variables from file
	if err := godotenv.Load(DefaultConfigFile); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Parse and load default arguments
	llamaCliArgs = ParseDefaultLlamaCliEnv()
	appArgs = ParseDefaultAppEnv()

	// Setup proper logging
	if err := setupLogging(); err != nil {
		log.Fatalf("Failed to setup logging: %v", err)
	}

	logger.Println("Application starting...")

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a separate goroutine
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- runServer(ctx)
	}()

	// Wait for shutdown signal or server error
	select {
	case <-quit:
		logger.Println("Received shutdown signal...")
	case err := <-serverErr:
		if err != nil && !errors.Is(err, context.Canceled) {
			logger.Printf("Server error: %v", err)
		}
	}

	// Cancel context and wait for graceful shutdown
	cancel()

	// Give server time to shutdown gracefully
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

// runServer starts and runs the MCP server
func runServer(ctx context.Context) error {
	// Create HTTP transport for MCP
	transport := mcphttp.NewHTTPTransport(appArgs.EndPoint)
	transport.WithAddr(appArgs.HttpPort)

	// Create MCP server
	server := mcpgolang.NewServer(transport)

	// Register the completion tool
	if err := server.RegisterTool(
		"generate_completion",
		"Generate text completion using the local LLM",
		handleCompletionTool,
	); err != nil {
		return fmt.Errorf("failed to register completion tool: %w", err)
	}

	logger.Printf("Starting MCP HTTP server on %s%s", appArgs.HttpPort, appArgs.EndPoint)

	// Start server with context cancellation
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Serve()
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.Done():
		logger.Println("Shutting down server...")
		if err := transport.Close(); err != nil {
			logger.Printf("Transport shutdown error: %v", err)
		}
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

// handleCompletionTool handles the completion tool requests
func handleCompletionTool(arguments CompletionArguments) (*mcpgolang.ToolResponse, error) {
	var metrics CompletionMetrics
	startTime := time.Now()
	metrics.RequestCount++

	defer func() {
		duration := time.Since(startTime)
		metrics.TotalDuration += duration
		logger.Printf("Request completed in %v (avg: %v)",
			duration,
			time.Duration(int64(metrics.TotalDuration)/metrics.RequestCount))
	}()

	// Simple check for empty prompt
	if arguments.Prompt == "" {
		logger.Println("Empty prompt received")
		return &mcpgolang.ToolResponse{
			Content: []*mcpgolang.Content{
				mcpgolang.NewTextContent("Error: Prompt cannot be empty"),
			},
		}, nil
	}

	logger.Printf("Handling completion request for prompt: %.100s...", arguments.Prompt)

	// Use timeout from configuration
	timeoutSeconds := appArgs.TimeOutSeconds
	if timeoutSeconds <= 0 {
		timeoutSeconds = 300 // fallback default
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	logger.Printf("Starting completion with timeout of %d seconds", timeoutSeconds)

	// Prepare arguments for llama using configuration from .env
	args := prepareLlamaArgs(arguments.Prompt)

	// Generate completion
	output, err := GenerateSingleCompletionWithCancel(ctx, appArgs, args)
	if err != nil {
		// Check if the error is due to timeout
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Printf("Completion timed out after %d seconds", timeoutSeconds)
			return &mcpgolang.ToolResponse{
				Content: []*mcpgolang.Content{
					mcpgolang.NewTextContent(fmt.Sprintf("Error: Completion timed out after %d seconds", timeoutSeconds)),
				},
			}, nil
		}

		logger.Printf("Error generating completion: %v", err)
		return &mcpgolang.ToolResponse{
			Content: []*mcpgolang.Content{
				mcpgolang.NewTextContent(fmt.Sprintf("Error generating completion: %v", err)),
			},
		}, nil
	}

	logger.Printf("Completion generated successfully, output length: %d chars", len(output))

	// Return the completion as tool response
	return &mcpgolang.ToolResponse{
		Content: []*mcpgolang.Content{
			mcpgolang.NewTextContent(string(output)),
		},
	}, nil
}

// prepareLlamaArgs prepares arguments for llamacpp using only the prompt and .env configuration
func prepareLlamaArgs(prompt string) []string {
	// Start with the base arguments from LlamaCliArgs (from .env)
	args := LlamaCliStructToArgs(llamaCliArgs)

	// Override the prompt with the one from MCP input
	// Remove any existing --prompt argument and add the new one
	filteredArgs := make([]string, 0, len(args))
	skipNext := false

	for _, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}
		if arg == "--prompt" {
			skipNext = true // skip the next argument (the old prompt value)
			continue
		}
		filteredArgs = append(filteredArgs, arg)
	}

	// Add the new prompt
	filteredArgs = append(filteredArgs, "--prompt", prompt)

	logger.Printf("Prepared llama args with %d parameters", len(filteredArgs))
	return filteredArgs
}
