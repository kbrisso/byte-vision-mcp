// Package main contains type definitions and configuration parsing for the
// Byte Vision MCP server, including LLama.cpp argument structures and
// environment variable parsing functions.
package main

import (
	"os"
	"strconv"
)

// ParseDefaultLlamaCliEnv parses all LLama.cpp related environment variables
// and returns a populated LlamaCliArgs struct with all command-line options.
//
// Returns:
//   - LlamaCliArgs: Struct containing all parsed LLama.cpp configuration
func ParseDefaultLlamaCliEnv() LlamaCliArgs {
	out := LlamaCliArgs{
		// Model configuration
		ModelCmd:         os.Getenv("ModelCmd"),
		ModelFullPathVal: os.Getenv("ModelFullPathVal"),

		// Prompt configuration
		PromptCmd:        os.Getenv("PromptCmd"),
		PromptCmdEnabled: getEnvBool(os.Getenv("PromptCmdEnabled"), false),
		PromptText:       os.Getenv("PromptText"),

		// Chat and input configuration
		ChatTemplateCmd:          os.Getenv("ChatTemplateCmd"),
		ChatTemplateVal:          os.Getenv("ChatTemplateVal"),
		MultilineInputCmd:        os.Getenv("MultilineInputCmd"),
		MultilineInputCmdEnabled: getEnvBool(os.Getenv("MultilineInputCmdEnabled"), false),

		// Context and scaling configuration
		CtxSizeCmd:        os.Getenv("CtxSizeCmd"),
		CtxSizeVal:        os.Getenv("CtxSizeVal"),
		RopeScalingCmd:    os.Getenv("RopeScalingCmd"),
		RopeScalingCmdVal: os.Getenv("RopeScalingCmdVal"),
		RopeScaleCmd:      os.Getenv("RopeScaleCmd"),
		RopeScaleVal:      os.Getenv("RopeScaleVal"),

		// Caching configuration
		PromptCacheAllCmd: os.Getenv("PromptCacheAllCmd"),
		PromptCacheCmd:    os.Getenv("PromptCacheCmd"),
		PromptCacheVal:    os.Getenv("PromptCacheVal"),

		// File and prompt handling
		PromptFileCmd:    os.Getenv("PromptFileCmd"),
		PromptFileVal:    os.Getenv("PromptFileVal"),
		ReversePromptCmd: os.Getenv("ReversePromptCmd"),
		ReversePromptVal: os.Getenv("ReversePromptVal"),
		InPrefixCmd:      os.Getenv("InPrefixCmd"),
		InPrefixVal:      os.Getenv("InPrefixVal"),
		InSuffixCmd:      os.Getenv("InSuffixCmd"),
		InSuffixVal:      os.Getenv("InSuffixVal"),

		// GPU and threading configuration
		GPULayersCmd:    os.Getenv("GPULayersCmd"),
		GPULayersVal:    os.Getenv("GPULayersVal"),
		ThreadsBatchCmd: os.Getenv("ThreadsBatchCmd"),
		ThreadsBatchVal: os.Getenv("ThreadsBatchVal"),
		ThreadsCmd:      os.Getenv("ThreadsCmd"),
		ThreadsVal:      os.Getenv("ThreadsVal"),

		// Generation parameters
		KeepCmd:              os.Getenv("KeepCmd"),
		KeepVal:              os.Getenv("KeepVal"),
		TopKCmd:              os.Getenv("TopKCmd"),
		TopKVal:              os.Getenv("TopKVal"),
		MainGPUCmd:           os.Getenv("MainGPUCmd"),
		MainGPUVal:           os.Getenv("MainGPUVal"),
		RepeatPenaltyCmd:     os.Getenv("RepeatPenaltyCmd"),
		RepeatPenaltyVal:     os.Getenv("RepeatPenaltyVal"),
		RepeatLastPenaltyCmd: os.Getenv("RepeatLastPenaltyCmd"),
		RepeatLastPenaltyVal: os.Getenv("RepeatLastPenaltyVal"),

		// Memory and system configuration
		MemLockCmd:               os.Getenv("MemLockCmd"),
		MemLockCmdEnabled:        getEnvBool(os.Getenv("MemLockCmdEnabled"), false),
		EscapeNewLinesCmd:        os.Getenv("EscapeNewLinesCmd"),
		EscapeNewLinesCmdEnabled: getEnvBool(os.Getenv("EscapeNewLinesCmdEnabled"), false),

		// Logging and debugging
		LogVerboseCmd:     os.Getenv("LogVerboseCmd"),
		LogVerboseEnabled: getEnvBool(os.Getenv("LogVerboseEnabled"), false),

		// Sampling parameters
		TemperatureVal: os.Getenv("TemperatureVal"),
		TemperatureCmd: os.Getenv("TemperatureCmd"),
		PredictCmd:     os.Getenv("PredictCmd"),
		PredictVal:     os.Getenv("PredictVal"),

		// Display and output configuration
		NoDisplayPromptCmd:     os.Getenv("NoDisplayPromptCmd"),
		NoDisplayPromptEnabled: getEnvBool(os.Getenv("NoDisplayPromptEnabled"), false),
		TopPCmd:                os.Getenv("TopPCmd"),
		TopPVal:                os.Getenv("TopPVal"),
		MinPCmd:                os.Getenv("MinPCmd"),
		MinPVal:                os.Getenv("MinPVal"),

		// Logging configuration
		ModelLogFileCmd:     os.Getenv("ModelLogFileCmd"),
		ModelLogFileNameVal: os.Getenv("ModelLogFileNameVal"),

		// Advanced features
		FlashAttentionCmd:        os.Getenv("FlashAttentionCmd"),
		FlashAttentionCmdEnabled: getEnvBool(os.Getenv("FlashAttentionCmdEnabled"), false),
		NoConversationCmd:        os.Getenv("NoConversationCmd"),
		NoConversationCmdEnabled: getEnvBool(os.Getenv("NoConversationCmdEnabled"), false),
		NoContextShiftCmd:        os.Getenv("NoContextShiftCmd"),
		NoContextShiftCmdEnabled: getEnvBool(os.Getenv("NoContextShiftCmdEnabled"), false),

		// Advanced parameters
		RandomSeedCmd:         os.Getenv("RandomSeedCmd"),
		RandomSeedCmdVal:      os.Getenv("RandomSeedCmdVal"),
		YarnOrigContextCmd:    os.Getenv("YarnOrigContextCmd"),
		YarnOrigContextCmdVal: os.Getenv("YarnOrigContextCmdVal"),

		// Batch processing configuration
		BatchCmd:     os.Getenv("BatchCmd"),
		BatchCmdVal:  os.Getenv("BatchCmdVal"),
		UBatchCmd:    os.Getenv("UBatchCmd"),
		UBatchCmdVal: os.Getenv("UBatchCmdVal"),

		// Model splitting configuration
		SplitModeCmd:    os.Getenv("SplitModeCmd"),
		SplitModeCmdVal: os.Getenv("SplitModeCmdVal"),
	}
	return out
}

// ParseDefaultAppEnv parses application-level environment variables
// and returns a populated DefaultAppArgs struct with general app configuration.
//
// Returns:
//   - DefaultAppArgs: Struct containing all parsed application configuration
func ParseDefaultAppEnv() DefaultAppArgs {
	out := DefaultAppArgs{
		// Path configurations
		ModelPath:       os.Getenv("ModelPath"),
		AppLogPath:      os.Getenv("AppLogPath"),
		AppLogFileName:  os.Getenv("AppLogFileName"),
		LLamaCliPath:    os.Getenv("LLamaCliPath"),
		PromptCachePath: os.Getenv("PromptCachePath"),

		// Server configuration
		HttpPort:       os.Getenv("HttpPort"),
		EndPoint:       os.Getenv("EndPoint"),
		TimeOutSeconds: getEnvInt("TimeOutSeconds", 300),
	}
	return out
}

// getEnvInt parses an environment variable as an integer with a fallback value.
// If the environment variable is empty or cannot be parsed, returns the fallback.
//
// Parameters:
//   - key: The environment variable name to parse
//   - fallback: The default value to return if parsing fails
//
// Returns:
//   - int: The parsed integer value or fallback
func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return fallback
}

// getEnvBool parses an environment variable as a boolean with a fallback value.
// Accepts standard boolean representations: "true", "false", "1", "0", etc.
//
// Parameters:
//   - key: The environment variable value to parse (not the key name)
//   - fallback: The default value to return if parsing fails
//
// Returns:
//   - bool: The parsed boolean value or fallback
func getEnvBool(key string, fallback bool) bool {
	// Return fallback if the value is empty
	if key == "" {
		return fallback
	}

	// Parse the boolean value
	result, err := strconv.ParseBool(key)
	if err != nil {
		return fallback
	}

	return result
}

// LlamaCliArgs contains all possible command-line arguments and flags
// that can be passed to the llama-cli executable for model inference.
type LlamaCliArgs struct {
	// Basic prompt configuration
	PromptCmd        string `json:"PromptCmd"`        // Command flag for prompt input (--prompt)
	PromptCmdEnabled bool   `json:"PromptCmdEnabled"` // Whether to enable prompt command
	PromptText       string `json:"PromptText"`       // The actual prompt text content

	// Chat template configuration
	ChatTemplateCmd string `json:"ChatTemplateCmd"` // Command flag for chat template (--chat-template)
	ChatTemplateVal string `json:"ChatTemplateVal"` // Chat template format string

	// Input handling configuration
	MultilineInputCmd        string `json:"MultilineInputCmd"`        // Command flag for multiline input (--multiline-input)
	MultilineInputCmdEnabled bool   `json:"MultilineInputCmdEnabled"` // Whether to enable multiline input

	// Context window configuration
	CtxSizeCmd string `json:"CtxSizeCmd"` // Command flag for context size (--ctx-size)
	CtxSizeVal string `json:"CtxSizeVal"` // Context window size value

	// RoPE scaling configuration for extended context
	RopeScaleCmd string `json:"RopeScaleCmd"` // Command flag for RoPE scale (--rope-scale)
	RopeScaleVal string `json:"RopeScaleVal"` // RoPE scale factor value

	// Advanced RoPE scaling configuration
	RopeScalingCmd    string `json:"RopeScalingCmd"`    // Command flag for RoPE scaling type
	RopeScalingCmdVal string `json:"RopeScalingCmdVal"` // RoPE scaling type value

	// Prompt caching configuration
	PromptCacheAllCmd     string `json:"PromptCacheAllCmd"`     // Command flag for cache all prompts
	PromptCacheAllEnabled bool   `json:"PromptCacheAllEnabled"` // Whether to enable cache all prompts
	PromptCacheCmd        string `json:"PromptCacheCmd"`        // Command flag for prompt cache (--prompt-cache)
	PromptCacheVal        string `json:"PromptCacheVal"`        // Prompt cache file path

	// File input configuration
	PromptFileCmd string `json:"PromptFileCmd"` // Command flag for prompt file input (--file)
	PromptFileVal string `json:"PromptFileVal"` // Prompt file path

	// Reverse prompt configuration for interactive mode
	ReversePromptCmd string `json:"ReversePromptCmd"` // Command flag for reverse prompt (--reverse-prompt)
	ReversePromptVal string `json:"ReversePromptVal"` // Reverse prompt text

	// Input formatting configuration
	InPrefixCmd string `json:"InPrefixCmd"` // Command flag for input prefix (--in-prefix)
	InPrefixVal string `json:"InPrefixVal"` // Input prefix text
	InSuffixCmd string `json:"InSuffixCmd"` // Command flag for input suffix (--in-suffix)
	InSuffixVal string `json:"InSuffixVal"` // Input suffix text

	// GPU acceleration configuration
	GPULayersCmd string `json:"GPULayersCmd"` // Command flag for GPU layers (--n-gpu-layers)
	GPULayersVal string `json:"GPULayersVal"` // Number of layers to offload to GPU

	// Threading configuration
	ThreadsBatchCmd string `json:"ThreadsBatchCmd"` // Command flag for batch threads (--threads-batch)
	ThreadsBatchVal string `json:"ThreadsBatchVal"` // Number of threads for batch processing
	ThreadsCmd      string `json:"ThreadsCmd"`      // Command flag for threads (--threads)
	ThreadsVal      string `json:"ThreadsVal"`      // Number of threads for inference

	// Context management configuration
	KeepCmd string `json:"KeepCmd"` // Command flag for keep tokens (--keep)
	KeepVal string `json:"KeepVal"` // Number of tokens to keep in context

	// Sampling parameters
	TopKCmd string `json:"TopKCmd"` // Command flag for top-k sampling (--top-k)
	TopKVal string `json:"TopKVal"` // Top-k sampling value

	// Multi-GPU configuration
	MainGPUCmd string `json:"MainGPUCmd"` // Command flag for main GPU (--main-gpu)
	MainGPUVal string `json:"MainGPUVal"` // Main GPU index

	// Repetition penalty configuration
	RepeatPenaltyCmd     string `json:"RepeatPenaltyCmd"`     // Command flag for repeat penalty (--repeat-penalty)
	RepeatPenaltyVal     string `json:"RepeatPenaltyVal"`     // Repeat penalty value
	RepeatLastPenaltyCmd string `json:"RepeatLastPenaltyCmd"` // Command flag for repeat last n (--repeat-last-n)
	RepeatLastPenaltyVal string `json:"RepeatLastPenaltyVal"` // Number of last tokens to consider for penalty

	// Memory management configuration
	MemLockCmd        string `json:"MemLockCmd"`        // Command flag for memory lock (--mlock)
	MemLockCmdEnabled bool   `json:"MemLockCmdEnabled"` // Whether to enable memory locking

	// Memory mapping configuration (unused in current implementation)
	NoMMApCmd        string `json:"NoMMApCmd"`        // Command flag for no memory mapping
	NoMMApCmdEnabled bool   `json:"NoMMApCmdEnabled"` // Whether to disable memory mapping

	// Output formatting configuration
	EscapeNewLinesCmd        string `json:"EscapeNewLinesCmd"`        // Command flag for escape newlines (-e)
	EscapeNewLinesCmdEnabled bool   `json:"EscapeNewLinesCmdEnabled"` // Whether to escape newlines in output

	// Debugging configuration
	LogVerboseCmd     string `json:"LogVerboseCmd"`     // Command flag for verbose logging (--log-verbose)
	LogVerboseEnabled bool   `json:"LogVerboseEnabled"` // Whether to enable verbose logging

	// Temperature sampling configuration
	TemperatureVal string `json:"TemperatureVal"` // Temperature value for sampling randomness
	TemperatureCmd string `json:"TemperatureCmd"` // Command flag for temperature (--temp)

	// Generation length configuration
	PredictCmd string `json:"PredictCmd"` // Command flag for prediction length (--n-predict)
	PredictVal string `json:"PredictVal"` // Maximum number of tokens to generate

	// Model configuration
	ModelFullPathVal string `json:"ModelFullPathVal"` // Full path to the model file
	ModelCmd         string `json:"ModelCmd"`         // Command flag for model (--model)

	// Display configuration
	NoDisplayPromptCmd     string `json:"NoDisplayPromptCmd"`     // Command flag for no display prompt (--no-display-prompt)
	NoDisplayPromptEnabled bool   `json:"NoDisplayPromptEnabled"` // Whether to hide prompt in output

	// Advanced sampling parameters
	TopPCmd string `json:"TopPCmd"` // Command flag for top-p sampling (--top-p)
	TopPVal string `json:"TopPVal"` // Top-p (nucleus) sampling value
	MinPCmd string `json:"MinPCmd"` // Command flag for min-p sampling (--min-p)
	MinPVal string `json:"MinPVal"` // Min-p sampling value

	// Model logging configuration
	ModelLogFileCmd     string `json:"ModelLogFileCmd"`     // Command flag for model log file (--log-file)
	ModelLogFileNameVal string `json:"ModelLogFileNameVal"` // Model log file path

	// Performance optimization features
	FlashAttentionCmd        string `json:"FlashAttentionCmd"`        // Command flag for flash attention (--flash-attn)
	FlashAttentionCmdEnabled bool   `json:"FlashAttentionCmdEnabled"` // Whether to enable flash attention

	// Conversation mode configuration
	NoConversationCmd        string `json:"NoConversationCmd"`        // Command flag for no conversation mode
	NoConversationCmdEnabled bool   `json:"NoConversationCmdEnabled"` // Whether to disable conversation mode

	// Context shifting configuration
	NoContextShiftCmd        string `json:"NoContextShiftCmd"`        // Command flag for no context shift
	NoContextShiftCmdEnabled bool   `json:"NoContextShiftCmdEnabled"` // Whether to disable context shifting

	// Random seed configuration
	RandomSeedCmd    string `json:"RandomSeedCmd"`    // Command flag for random seed (--seed)
	RandomSeedCmdVal string `json:"RandomSeedCmdVal"` // Random seed value for reproducible generation

	// YaRN (Yet another RoPE extensioN) configuration
	YarnOrigContextCmd    string `json:"YarnOrigContextCmd"`    // Command flag for YaRN original context
	YarnOrigContextCmdVal string `json:"YarnOrigContextCmdVal"` // YaRN original context size

	// Batch processing configuration
	BatchCmd     string `json:"BatchCmd"`     // Command flag for batch size (--batch-size)
	BatchCmdVal  string `json:"BatchCmdVal"`  // Batch size for processing
	UBatchCmd    string `json:"UBatchCmd"`    // Command flag for micro-batch size (--ubatch-size)
	UBatchCmdVal string `json:"UBatchCmdVal"` // Micro-batch size for processing

	// Model splitting configuration for multi-GPU setups
	SplitModeCmd    string `json:"SplitModeCmd"`    // Command flag for split mode (--split-mode)
	SplitModeCmdVal string `json:"SplitModeCmdVal"` // Split mode value (layer, row, etc.)
	Prompt          string `json:"prompt" description:"The prompt text to generate completion for"`

	// Core Model & Performance Parameters
	Model     string `json:"model,omitempty" description:"Model path (overrides default)"`
	Threads   int    `json:"threads,omitempty" description:"CPU threads for generation"`
	GpuLayers int    `json:"gpu_layers,omitempty" description:"GPU acceleration layers"`
	CtxSize   int    `json:"ctx_size,omitempty" description:"Context window size"`
	BatchSize int    `json:"batch_size,omitempty" description:"Batch processing size"`

	// Generation Control Parameters
	Predict       int     `json:"predict,omitempty" description:"Number of tokens to generate"`
	Temperature   float64 `json:"temperature,omitempty" description:"Creativity/randomness control"`
	TopK          int     `json:"top_k,omitempty" description:"Top-K sampling"`
	TopP          float64 `json:"top_p,omitempty" description:"Top-P (nucleus) sampling"`
	RepeatPenalty float64 `json:"repeat_penalty,omitempty" description:"Repetition penalty"`

	// Input/Output Parameters
	PromptFile string `json:"prompt_file,omitempty" description:"Prompt from file"`
	LogFile    string `json:"log_file,omitempty" description:"Output logging"`
}

// DefaultAppArgs contains general application configuration parameters
// that are not specific to LLama.cpp but control the MCP server behavior.
type DefaultAppArgs struct {
	ModelPath       string `json:"ModelPath"`       // Directory path where model files are stored
	AppLogPath      string `json:"AppLogPath"`      // Directory path for application log files
	AppLogFileName  string `json:"AppLogFileName"`  // Name of the main application log file
	PromptCachePath string `json:"PromptCachePath"` // Directory path for prompt cache files
	LLamaCliPath    string `json:"LlamaCliPath"`    // Full path to the llama-cli executable
	HttpPort        string `json:"HttpPort"`        // HTTP port for the MCP server (e.g., ":8080")
	EndPoint        string `json:"EndPoint"`        // HTTP endpoint path for MCP requests (e.g., "/mcp-completion")
	TimeOutSeconds  int    `json:"TimeOutSeconds"`  // Timeout in seconds for completion requests
}
