# Byte Vision MCP

A Model Context Protocol (MCP) server that provides text completion capabilities using local LLama.cpp models. This
server exposes a single MCP tool that accepts text prompts and returns AI-generated completions using locally hosted
language models.

## What is this project?

Byte Vision MCP is a bridge between MCP-compatible clients (like Claude Desktop, IDEs, or other AI tools) and local
LLama.cpp language models. It allows you to:

- Use local language models through the MCP protocol
- Configure all model parameters via environment files
- Generate text completions with custom prompts
- Maintain privacy by keeping everything local
- Integrate with MCP-compatible applications

## Features

- **MCP Protocol Support**: Standard MCP server implementation
- **Local Model Execution**: Uses LLama.cpp for model inference
- **Configurable Parameters**: All settings controlled via environment file
- **GPU Acceleration**: Supports CUDA, ROCm, and Metal
- **Prompt Caching**: Built-in caching for improved performance
- **Comprehensive Logging**: Detailed logging for debugging and monitoring
- **Graceful Shutdown**: Proper resource cleanup and error handling

## Built With

### Core Dependencies

- **[github.com/joho/godotenv](https://github.com/joho/godotenv) v1.5.1** - Environment variable loading from `.env`
  files
- **[github.com/metoro-io/mcp-golang](https://github.com/metoro-io/mcp-golang) v0.12.0** - Model Context Protocol (MCP)
  implementation for Go

### Indirect Dependencies

#### Web Framework & HTTP

- **[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) v1.8.1** - HTTP web framework
- **[github.com/gin-contrib/sse](https://github.com/gin-contrib/sse) v0.1.0** - Server-Sent Events support

### Runtime Requirements

- **Go SDK 1.23+** - Modern Go runtime with latest features
- **LLama.cpp** - Local language model inference engine
- **GGUF Models** - Quantized language models in GGUF format

## Prerequisites

- **Go 1.23+** for building the server
- **LLama.cpp binaries** (see `/llamacpp/README.md` for installation)
- **GGUF format models** (see `/models/README.md` for sources)

## Quick Start

### 1. Clone and Build

```bash
git clone <repository-url>
cd byte-vision-mcp
go mod tidy
go build -o byte-vision-mcp
```

### 2. Set Up LLama.cpp

Follow the instructions in: `/llamacpp/README.md`

- Download prebuilt binaries, or
- Build from source

### 3. Download Models

See: `/models/README.md`

- Recommended model sources
- How to download GGUF models
- Model placement instructions

### 4. Configure Environment

Copy the example configuration:
```
cp example-byte-vision-cfg.env byte-vision-cfg.env
```
Edit to match your setup: `byte-vision-cfg.env`

# Update paths to match your installation

LLamaCliPath=/path/to/your/llama-cli
ModelFullPathVal=/path/to/your/model.gguf
AppLogPath=/path/to/logs/

### 5. Run the Server
```
./byte-vision-mcp
```
The server will start on `http://localhost:8080/mcp-completion` by default.

## Project Structure
```
byte-vision-mcp/
├── llamacpp/ # LLama.cpp binaries and installation guide
├── logs/ # Application and model logs
├── models/ # GGUF model files
├── prompt-cache/ # Cached prompts for performance
├── main.go # Main MCP server implementation
├── model.go # Model execution logic
├── types.go # Configuration structures
├── byte-vision-cfg.env # Your configuration (create from example)
└── example-byte-vision-cfg.env # Example configuration
```
## Configuration

The file controls all aspects of the server: `byte-vision-cfg.env`

### Application Settings

- : Directory for log files `AppLogPath`
- : Log file name `AppLogFileName`
- : Server port (default `:8080`) `HttpPort`
- : MCP endpoint path (default ) `EndPoint``/mcp-completion`
- : Request timeout (default `300`) `TimeOutSeconds`

### LLama.cpp Settings

- : Path to llama-cli executable `LLamaCliPath`
- : Path to your GGUF model file `ModelFullPathVal`
- : Context window size `CtxSizeVal`
- : Number of layers to offload to GPU `GPULayersVal`
- : Generation temperature `TemperatureVal`
- : Maximum tokens to generate `PredictVal`
- And many more LLama.cpp parameters...

## Usage

### MCP Tool: `generate_completion`

The server exposes a single MCP tool that accepts both basic and advanced parameters for fine-tuning LLama.cpp
execution.

#### Basic Usage
```json
**Input:**
{
"prompt": "Write a short story about a robot learning to paint"
}
**Output:**
{
"content": [
{
"type": "text",
"text": "Generated completion text..."
}
]
}
```
#### Advanced Usage with Parameters

**Full Parameter Input:**
```json
{
"prompt": "Explain quantum computing in simple terms",
"temperature": 0.7,
"predict": 500,
"top_k": 40,
"top_p": 0.9,
"ctx_size": 4096,
"threads": 8,
"gpu_layers": 35
}
```
#### Available Parameters

##### Core Model & Performance Parameters

| Parameter    | Type   | Description                | Example                 | Default Source     |
|--------------|--------|----------------------------|-------------------------|--------------------|
| `model`      | string | Override model path        | `"/path/to/model.gguf"` | `ModelFullPathVal` |
| `threads`    | int    | CPU threads for generation | `8`                     | `ThreadsVal`       |
| `gpu_layers` | int    | GPU acceleration layers    | `35`                    | `GPULayersVal`     |
| `ctx_size`   | int    | Context window size        | `4096`                  | `CtxSizeVal`       |
| `batch_size` | int    | Batch processing size      | `512`                   | `BatchCmdVal`      |

##### Generation Control Parameters

| Parameter        | Type  | Description                   | Range     | Default Source     |
|------------------|-------|-------------------------------|-----------|--------------------|
| `predict`        | int   | Number of tokens to generate  | `1-8192`  | `PredictVal`       |
| `temperature`    | float | Creativity/randomness control | `0.0-2.0` | `TemperatureVal`   |
| `top_k`          | int   | Top-K sampling                | `1-100`   | `TopKVal`          |
| `top_p`          | float | Top-P (nucleus) sampling      | `0.0-1.0` | `TopPVal`          |
| `repeat_penalty` | float | Repetition penalty            | `0.5-2.0` | `RepeatPenaltyVal` |

##### Input/Output Parameters

| Parameter     | Type   | Description           | Example                 | Default Source        |
|---------------|--------|-----------------------|-------------------------|-----------------------|
| `prompt_file` | string | Load prompt from file | `"/path/to/prompt.txt"` | `PromptFileVal`       |
| `log_file`    | string | Custom log file path  | `"/path/to/custom.log"` | `ModelLogFileNameVal` |

#### Parameter Usage Examples

##### 1. Creative Writing (High Temperature)
```json
{
"prompt": "Write a creative story about time travel",
"temperature": 1.2,
"top_p": 0.95,
"predict": 1000,
"repeat_penalty": 1.1
}
```

##### 2. Technical Documentation (Low Temperature)
```json
{
"prompt": "Explain the TCP/IP protocol stack",
"temperature": 0.3,
"top_k": 10,
"predict": 800,
"ctx_size": 8192
}
```
##### 3. Code Generation (Balanced)
```json
{
"prompt": "Write a Python function to sort a list",
"temperature": 0.6,
"top_k": 30,
"top_p": 0.8,
"predict": 400
}
```
##### 4. Long Context Processing
```json
{
"prompt": "Summarize this document...",
"ctx_size": 32768,
"gpu_layers": 40,
"batch_size": 1024,
"predict": 500
}
```
##### 5. Performance Optimization
```json
{
"prompt": "Quick question about Go syntax",
"threads": 12,
"gpu_layers": 45,
"batch_size": 2048,
"predict": 200,
"temperature": 0.4
}
```
##### 6. Using External Prompt File
```json
{
"prompt_file": "/path/to/complex_prompt.txt",
"temperature": 0.8,
"predict": 1500,
"log_file": "/path/to/custom_generation.log"
}
```
#### Parameter Guidelines

##### Temperature Settings
```
- **0.0-0.3**: Highly deterministic, factual responses
- **0.4-0.7**: Balanced creativity and coherence
- **0.8-1.2**: Creative, varied responses
- **1.3-2.0**: Highly creative, potentially chaotic
```
##### Context Size Guidelines
```
- **2048-4096**: Short conversations, simple tasks
- **8192-16384**: Medium documents, complex reasoning
- **32768+**: Long documents, extensive context
```
##### GPU Layers Optimization
```
- **0**: CPU-only processing
- **25-35**: Balanced CPU/GPU (8GB VRAM)
- **40+**: Full GPU acceleration (12GB+ VRAM)
```
##### Prediction Length
```
- **50-200**: Short answers, code snippets
- **300-800**: Medium explanations, documentation
- **1000+**: Long-form content, stories
```
#### Performance Considerations

##### Memory Usage
```json
{
"ctx_size": 4096, // Lower for limited RAM
"batch_size": 512, // Smaller batches for stability
"gpu_layers": 25 // Reduce if GPU memory limited
}
```
##### Speed Optimization
```json
{
"threads": 8, // Match CPU cores
"gpu_layers": 45, // Maximize GPU usage
"batch_size": 2048, // Larger batches for throughput
"predict": 200 // Shorter for quick responses
}
```
##### Quality Optimization
```json
{
"temperature": 0.7, // Balanced creativity
"top_k": 40, // Diverse sampling
"top_p": 0.9, // Nucleus sampling
"repeat_penalty": 1.1, // Reduce repetition
"ctx_size": 8192 // Larger context for coherence
}
```
#### Error Handling

The tool provides specific error messages for invalid parameters:
```json
{
"content": [
{
"type": "text",
"text": "Error: Invalid temperature value. Must be between 0.0 and 2.0"
}
]
}
```
#### Default Behavior

- **All parameters are optional** except `prompt`
- **Environment configuration** is used when parameters are not specified
- **Zero values are ignored** (e.g., `temperature: 0` uses config default)
- **Invalid values** fall back to configuration defaults

#### Integration Examples

##### JavaScript/Node.js
```javascript
const result = await mcpClient.callTool("generate_completion", {
prompt: "Explain async/await in JavaScript",
temperature: 0.5,
predict: 600,
top_k: 25
});
console.log(result.content[0].text);
```
##### Python
```python
response = mcp_client.call_tool("generate_completion", {
"prompt": "Write a Python class for a binary tree",
"temperature": 0.6,
"predict": 800,
"top_p": 0.8
})
print(response["content"][0]["text"])
```
##### curl (Direct HTTP)
```curl
curl -X POST http://localhost:8080/mcp-completion \
-H "Content-Type: application/json" \
-d '{
"prompt": "Explain Docker containers",
"temperature": 0.4,
"predict": 500,
"ctx_size": 4096
}'
```
This comprehensive parameter system allows fine-grained control over LLama.cpp behavior while maintaining backward
compatibility and ease of use.

## GPU Acceleration

### NVIDIA GPUs (CUDA)

- Download CUDA-enabled LLama.cpp binaries
- Set `GPULayersVal=33` (or adjust based on your GPU memory)
- Set `MainGPUVal=0` (or your preferred GPU index)

### AMD GPUs (ROCm - Linux only)

- Download ROCm-enabled LLama.cpp binaries
- Configure similar to CUDA setup

### Apple Silicon (Metal - macOS)

- Metal support is built-in
- No additional configuration needed

## Logging

Logs are written to both console and file:

- Application logs: `logs/byte-vision-mcp.log`
- Model logs: `logs/[model-name].log`
- Configurable log levels and verbosity

See for log management details. `/logs/README.md`

## Troubleshooting

### Common Issues

1. **"llama-cli not found"**
    - Check in your file `LLamaCliPath``.env`
    - Ensure the binary has execute permissions

2. **"Model file not found"**
    - Verify points to a valid `.gguf` file `ModelFullPathVal`
    - Check file permissions

3. **Out of memory errors**
    - Reduce `CtxSizeVal`
    - Use a smaller model
    - Increase for GPU offloading `GPULayersVal`

4. **Slow generation**
    - Enable GPU acceleration
    - Increase `GPULayersVal`
    - Use quantized models (Q4, Q5, Q8)

5. **Server won't start**
    - Check if port is already in use
    - Verify all paths in configuration exist
    - Check logs for detailed error messages

## Development

### Building from Source

go mod tidy
go build -o byte-vision-mcp

### Running Tests

go test ./...

### Dependencies

-
    - Environment file loading `github.com/joho/godotenv`
-
    - MCP protocol implementation `github.com/metoro-io/mcp-golang`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the terms of the MIT license.

## Support

- Check the individual README files in each subdirectory for specific setup instructions
- Review logs for detailed error information
- Ensure all paths in configuration are absolute and accessible

**Author:** Kevin Brisson
**Email:** kbrisso@gmail.com
**LinkedIn:** [Kevin Brisson](https://linkedin.com/in/kevinbrisson)
**Project Link:** [https://github.com/kbrisso/byte-vision-mcp](https://github.com/kbrisso/byte-vision-mcp)
