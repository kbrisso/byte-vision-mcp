package main

import (
	"os"
	"strconv"
)

func ParseDefaultLlamaCliEnv() LlamaCliArgs {
	out := LlamaCliArgs{
		ModelCmd:                 os.Getenv("ModelCmd"),
		ModelFullPathVal:         os.Getenv("ModelFullPathVal"),
		PromptCmd:                os.Getenv("PromptCmd"),
		PromptCmdEnabled:         getEnvBool(os.Getenv("PromptCmdEnabled"), false),
		PromptText:               os.Getenv("PromptText"),
		ChatTemplateCmd:          os.Getenv("ChatTemplateCmd"),
		ChatTemplateVal:          os.Getenv("ChatTemplateVal"),
		MultilineInputCmd:        os.Getenv("MultilineInputCmd"),
		MultilineInputCmdEnabled: getEnvBool(os.Getenv("MultilineInputCmdEnabled"), false),
		CtxSizeCmd:               os.Getenv("CtxSizeCmd"),
		CtxSizeVal:               os.Getenv("CtxSizeVal"),
		RopeScalingCmd:           os.Getenv("RopeScalingCmd"),
		RopeScalingCmdVal:        os.Getenv("RopeScalingCmdVal"),
		RopeScaleCmd:             os.Getenv("RopeScaleCmd"),
		RopeScaleVal:             os.Getenv("RopeScaleVal"),
		PromptCacheAllCmd:        os.Getenv("PromptCacheAllCmd"),
		PromptCacheCmd:           os.Getenv("PromptCacheCmd"),
		PromptCacheVal:           os.Getenv("PromptCacheVal"),
		PromptFileCmd:            os.Getenv("PromptFileCmd"),
		PromptFileVal:            os.Getenv("PromptFileVal"),
		ReversePromptCmd:         os.Getenv("ReversePromptCmd"),
		ReversePromptVal:         os.Getenv("ReversePromptVal"),
		InPrefixCmd:              os.Getenv("InPrefixCmd"),
		InPrefixVal:              os.Getenv("InPrefixVal"),
		InSuffixCmd:              os.Getenv("InSuffixCmd"),
		InSuffixVal:              os.Getenv("InSuffixVal"),
		GPULayersCmd:             os.Getenv("GPULayersCmd"),
		GPULayersVal:             os.Getenv("GPULayersVal"),
		ThreadsBatchCmd:          os.Getenv("ThreadsBatchCmd"),
		ThreadsBatchVal:          os.Getenv("ThreadsBatchVal"),
		ThreadsCmd:               os.Getenv("ThreadsCmd"),
		ThreadsVal:               os.Getenv("ThreadsVal"),
		KeepCmd:                  os.Getenv("KeepCmd"),
		KeepVal:                  os.Getenv("KeepVal"),
		TopKCmd:                  os.Getenv("TopKCmd"),
		TopKVal:                  os.Getenv("TopKVal"),
		MainGPUCmd:               os.Getenv("MainGPUCmd"),
		MainGPUVal:               os.Getenv("MainGPUVal"),
		RepeatPenaltyCmd:         os.Getenv("RepeatPenaltyCmd"),
		RepeatPenaltyVal:         os.Getenv("RepeatPenaltyVal"),
		RepeatLastPenaltyCmd:     os.Getenv("RepeatLastPenaltyCmd"),
		RepeatLastPenaltyVal:     os.Getenv("RepeatLastPenaltyVal"),
		MemLockCmd:               os.Getenv("MemLockCmd"),
		MemLockCmdEnabled:        getEnvBool(os.Getenv("MemLockCmdEnabled"), false),
		EscapeNewLinesCmd:        os.Getenv("EscapeNewLinesCmd"),
		EscapeNewLinesCmdEnabled: getEnvBool(os.Getenv("EscapeNewLinesCmdEnabled"), false),
		LogVerboseCmd:            os.Getenv("LogVerboseCmd"),
		LogVerboseEnabled:        getEnvBool(os.Getenv("LogVerboseEnabled"), false),
		TemperatureVal:           os.Getenv("TemperatureVal"),
		TemperatureCmd:           os.Getenv("TemperatureCmd"),
		PredictCmd:               os.Getenv("PredictCmd"),
		PredictVal:               os.Getenv("PredictVal"),
		NoDisplayPromptCmd:       os.Getenv("NoDisplayPromptCmd"),
		NoDisplayPromptEnabled:   getEnvBool(os.Getenv("NoDisplayPromptEnabled"), false),
		TopPCmd:                  os.Getenv("TopPCmd"),
		TopPVal:                  os.Getenv("TopPVal"),
		MinPCmd:                  os.Getenv("MinPCmd"),
		MinPVal:                  os.Getenv("MinPVal"),
		ModelLogFileCmd:          os.Getenv("ModelLogFileCmd"),
		ModelLogFileNameVal:      os.Getenv("ModelLogFileNameVal"),
		FlashAttentionCmd:        os.Getenv("FlashAttentionCmd"),
		FlashAttentionCmdEnabled: getEnvBool(os.Getenv("FlashAttentionCmdEnabled"), false),
		NoConversationCmd:        os.Getenv("NoConversationCmd"),
		NoConversationCmdEnabled: getEnvBool(os.Getenv("NoConversationCmdEnabled"), false),
		NoContextShiftCmd:        os.Getenv("NoContextShiftCmd"),
		NoContextShiftCmdEnabled: getEnvBool(os.Getenv("NoContextShiftCmdEnabled"), false),
		RandomSeedCmd:            os.Getenv("RandomSeedCmd"),
		RandomSeedCmdVal:         os.Getenv("RandomSeedCmdVal"),
		YarnOrigContextCmd:       os.Getenv("YarnOrigContextCmd"),
		YarnOrigContextCmdVal:    os.Getenv("YarnOrigContextCmdVal"),
		BatchCmd:                 os.Getenv("BatchCmd"),
		BatchCmdVal:              os.Getenv("BatchCmdVal"),
		UBatchCmd:                os.Getenv("UBatchCmd"),
		UBatchCmdVal:             os.Getenv("UBatchCmdVal"),
		SplitModeCmd:             os.Getenv("SplitModeCmd"),
		SplitModeCmdVal:          os.Getenv("SplitModeCmdVal"),
	}
	return out
}

func ParseDefaultAppEnv() DefaultAppArgs {
	out := DefaultAppArgs{
		ModelPath:       os.Getenv("ModelPath"),
		AppLogPath:      os.Getenv("AppLogPath"),
		AppLogFileName:  os.Getenv("AppLogFileName"),
		LLamaCliPath:    os.Getenv("LLamaCliPath"),
		PromptCachePath: os.Getenv("PromptCachePath"),
		HttpPort:        os.Getenv("HttpPort"),
		EndPoint:        os.Getenv("EndPoint"),
		TimeOutSeconds:  getEnvInt("TimeOutSeconds", 300),
	}
	return out
}
func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return fallback
}
func getEnvBool(key string, fallback bool) bool {
	result, err := strconv.ParseBool(key)
	if err != nil {
		return fallback
	}

	return result
}

func LlamaCliStructToArgs(args LlamaCliArgs) []string {
	var result []string
	// Helper function for NullString pairs
	addCmdValPair := func(cmd string, val string) {
		if len(val) != 0 {
			result = append(result, cmd, val)
		}
	}
	// Helper function for NullString and NullBool pairs
	addCmdBoolPair := func(cmd string, val bool) {
		if val == true {
			// Convert boolean to string before adding it to the result
			result = append(result, cmd)
		}
	}
	// Examples of using the helper functions
	addCmdValPair(args.ModelCmd, args.ModelFullPathVal)
	addCmdValPair(args.PromptCmd, args.PromptText)
	addCmdValPair(args.ChatTemplateCmd, args.ChatTemplateVal)
	addCmdBoolPair(args.MultilineInputCmd, args.MultilineInputCmdEnabled)
	addCmdValPair(args.CtxSizeCmd, args.CtxSizeVal)
	addCmdValPair(args.RopeScaleCmd, args.RopeScaleVal)
	addCmdValPair(args.PromptCacheCmd, args.PromptCacheVal)
	addCmdValPair(args.PromptFileCmd, args.PromptFileVal)
	addCmdValPair(args.ReversePromptCmd, args.ReversePromptVal)
	addCmdValPair(args.InPrefixCmd, args.InPrefixVal)
	addCmdValPair(args.InSuffixCmd, args.InSuffixVal)
	addCmdValPair(args.GPULayersCmd, args.GPULayersVal)
	addCmdValPair(args.ThreadsBatchCmd, args.ThreadsBatchVal)
	addCmdValPair(args.ThreadsCmd, args.ThreadsVal)
	addCmdValPair(args.KeepCmd, args.KeepVal)
	addCmdValPair(args.TopKCmd, args.TopKVal)
	addCmdValPair(args.MainGPUCmd, args.MainGPUVal)
	addCmdValPair(args.RepeatPenaltyCmd, args.RepeatPenaltyVal)
	addCmdValPair(args.RepeatLastPenaltyCmd, args.RepeatLastPenaltyVal)
	addCmdBoolPair(args.MemLockCmd, args.MemLockCmdEnabled)
	addCmdBoolPair(args.EscapeNewLinesCmd, args.EscapeNewLinesCmdEnabled)
	addCmdValPair(args.TemperatureCmd, args.TemperatureVal)
	addCmdValPair(args.PredictCmd, args.PredictVal)
	addCmdValPair(args.ModelLogFileCmd, args.ModelLogFileNameVal)
	addCmdBoolPair(args.NoDisplayPromptCmd, args.NoDisplayPromptEnabled)
	addCmdValPair(args.TopPCmd, args.TopPVal)
	addCmdValPair(args.MinPCmd, args.MinPVal)
	addCmdBoolPair(args.LogVerboseCmd, args.LogVerboseEnabled)
	addCmdBoolPair(args.FlashAttentionCmd, args.FlashAttentionCmdEnabled)
	addCmdBoolPair(args.NoConversationCmd, args.NoConversationCmdEnabled)
	addCmdBoolPair(args.NoContextShiftCmd, args.NoContextShiftCmdEnabled)
	addCmdValPair(args.RopeScalingCmd, args.RopeScalingCmdVal)
	addCmdValPair(args.RandomSeedCmd, args.RandomSeedCmdVal)
	addCmdValPair(args.YarnOrigContextCmd, args.YarnOrigContextCmdVal)
	addCmdValPair(args.BatchCmd, args.BatchCmdVal)
	addCmdValPair(args.UBatchCmd, args.UBatchCmdVal)
	addCmdValPair(args.SplitModeCmd, args.SplitModeCmdVal)
	return result

}

type LlamaCliArgs struct {
	PromptCmd                string `json:"PromptCmd"`
	PromptCmdEnabled         bool   `json:"PromptCmdEnabled"`
	ChatTemplateCmd          string `json:"ChatTemplateCmd"`
	ChatTemplateVal          string `json:"ChatTemplateVal"`
	MultilineInputCmd        string `json:"MultilineInputCmd"`
	MultilineInputCmdEnabled bool   `json:"MultilineInputCmdEnabled"`
	CtxSizeCmd               string `json:"CtxSizeCmd"`
	CtxSizeVal               string `json:"CtxSizeVal"`
	RopeScaleVal             string `json:"RopeScaleVal"`
	RopeScaleCmd             string `json:"RopeScaleCmd"`
	PromptCacheAllCmd        string `json:"PromptCacheAllCmd"`
	PromptCacheAllEnabled    bool   `json:"PromptCacheAllEnabled"`
	PromptCacheCmd           string `json:"PromptCacheCmd"`
	PromptCacheVal           string `json:"PromptCacheVal"`
	PromptFileCmd            string `json:"PromptFileCmd"`
	PromptFileVal            string `json:"PromptFileVal"`
	ReversePromptCmd         string `json:"ReversePromptCmd"`
	ReversePromptVal         string `json:"ReversePromptVal"`
	InPrefixCmd              string `json:"InPrefixCmd"`
	InPrefixVal              string `json:"InPrefixVal"`
	InSuffixCmd              string `json:"InSuffixCmd"`
	InSuffixVal              string `json:"InSuffixVal"`
	GPULayersCmd             string `json:"GPULayersCmd"`
	GPULayersVal             string `json:"GPULayersVal"`
	ThreadsBatchCmd          string `json:"ThreadsBatchCmd"`
	ThreadsBatchVal          string `json:"ThreadsBatchVal"`
	ThreadsCmd               string `json:"ThreadsCmd"`
	ThreadsVal               string `json:"ThreadsVal"`
	KeepCmd                  string `json:"KeepCmd"`
	KeepVal                  string `json:"KeepVal"`
	TopKCmd                  string `json:"TopKCmd"`
	TopKVal                  string `json:"TopKVal"`
	MainGPUCmd               string `json:"MainGPUCmd"`
	MainGPUVal               string `json:"MainGPUVal"`
	RepeatPenaltyCmd         string `json:"RepeatPenaltyCmd"`
	RepeatPenaltyVal         string `json:"RepeatPenaltyVal"`
	RepeatLastPenaltyCmd     string `json:"RepeatLastPenaltyCmd"`
	RepeatLastPenaltyVal     string `json:"RepeatLastPenaltyVal"`
	MemLockCmd               string `json:"MemLockCmd"`
	MemLockCmdEnabled        bool   `json:"MemLockCmdEnabled"`
	NoMMApCmd                string `json:"NoMMApCmd"`
	NoMMApCmdEnabled         bool   `json:"NoMMApCmdEnabled"`
	EscapeNewLinesCmd        string `json:"EscapeNewLinesCmd"`
	EscapeNewLinesCmdEnabled bool   `json:"EscapeNewLinesCmdEnabled"`
	LogVerboseCmd            string `json:"LogVerboseCmd"`
	LogVerboseEnabled        bool   `json:"LogVerboseEnabled"`
	TemperatureVal           string `json:"TemperatureVal"`
	TemperatureCmd           string `json:"TemperatureCmd"`
	PredictCmd               string `json:"PredictCmd"`
	PredictVal               string `json:"PredictVal"`
	ModelFullPathVal         string `json:"ModelFullPathVal"`
	ModelCmd                 string `json:"ModelCmd"`
	PromptText               string `json:"PromptText"`
	NoDisplayPromptCmd       string `json:"NoDisplayPromptCmd"`
	NoDisplayPromptEnabled   bool   `json:"NoDisplayPromptEnabled"`
	TopPCmd                  string `json:"TopPCmd"`
	TopPVal                  string `json:"TopPVal"`
	MinPCmd                  string `json:"MinPCmd"`
	MinPVal                  string `json:"MinPVal"`
	ModelLogFileCmd          string `json:"ModelLogFileCmd"`
	ModelLogFileNameVal      string `json:"ModelLogFileNameVal"`
	FlashAttentionCmd        string `json:"FlashAttentionCmd"`
	FlashAttentionCmdEnabled bool   `json:"FlashAttentionCmdEnabled"`
	NoConversationCmd        string `json:"NoConversationCmd"`
	NoConversationCmdEnabled bool   `json:"NoConversationCmdEnabled"`
	NoContextShiftCmd        string `json:"NoContextShiftCmd"`
	NoContextShiftCmdEnabled bool   `json:"NoContextShiftCmdEnabled"`
	RopeScalingCmd           string `json:"RopeScalingCmd"`
	RopeScalingCmdVal        string `json:"RopeScalingCmdVal"`
	RandomSeedCmd            string `json:"RandomSeedCmd"`
	RandomSeedCmdVal         string `json:"RandomSeedCmdVal"`
	YarnOrigContextCmd       string `json:"YarnOrigContextCmd"`
	YarnOrigContextCmdVal    string `json:"YarnOrigContextCmdVal"`
	BatchCmd                 string `json:"BatchCmd"`
	BatchCmdVal              string `json:"BatchCmdVal"`
	UBatchCmd                string `json:"UBatchCmd"`
	UBatchCmdVal             string `json:"UBatchCmdVal"`
	SplitModeCmd             string `json:"SplitModeCmd"`
	SplitModeCmdVal          string `json:"SplitModeCmdVal"`
}

type DefaultAppArgs struct {
	ModelPath       string `json:"ModelPath"`
	AppLogPath      string `json:"AppLogPath"`
	AppLogFileName  string `json:"AppLogFileName"`
	PromptCachePath string `json:"PromptCachePath"`
	LLamaCliPath    string `json:"LlamaCliPath"`
	HttpPort        string `json:"HttpPort"`
	EndPoint        string `json:"EndPoint"`
	TimeOutSeconds  int    `json:"TimeOutSeconds"`
}
