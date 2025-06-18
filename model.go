package main

import (
	"context"
	"os/exec"
)

// GenerateSingleCompletionWithCancel executes a LLama.cpp command with cancellation support.
// It runs the command in a separate goroutine to allow for context cancellation and timeouts.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - appArgs: Application configuration containing the path to llama-cli
//   - args: Command-line arguments to pass to llama-cli
//
// Returns:
//   - []byte: The output from the LLama.cpp command
//   - error: Any error that occurred during execution or context cancellation
func GenerateSingleCompletionWithCancel(ctx context.Context, appArgs DefaultAppArgs, args []string) ([]byte, error) {
	// Create a child context with cancel to ensure proper cleanup
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a channel to capture the command execution result
	// Using an anonymous struct to bundle output and error together
	result := make(chan struct {
		output []byte
		err    error
	})

	// Execute the command in a separate goroutine to enable cancellation
	go func() {
		// Run llama-cli with the provided arguments and context
		out, err := exec.CommandContext(ctx, appArgs.LLamaCliPath, args...).Output()

		// Send the result back through the channel
		result <- struct {
			output []byte
			err    error
		}{output: out, err: err}

		// Close the channel to signal completion
		close(result)
	}()

	// Wait for either command completion or context cancellation
	select {
	case res := <-result:
		// Command completed successfully or with an error
		return res.output, res.err
	case <-ctx.Done():
		// Context was canceled or timed out
		return nil, ctx.Err()
	}
}
