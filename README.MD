# Byte Vision MCP

A Model Context Protocol (MCP) server that provides text completion capabilities using local LLama.cpp models. This server exposes a single MCP tool that accepts text prompts and returns AI-generated completions using locally hosted language models.

## What is this project?

Byte Vision MCP is a bridge between MCP-compatible clients (like Claude Desktop, IDEs, or other AI tools) and local LLama.cpp language models. It allows you to:

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

- **[github.com/joho/godotenv](https://github.com/joho/godotenv) v1.5.1** - Environment variable loading from `.env` files
- **[github.com/metoro-io/mcp-golang](https://github.com/metoro-io/mcp-golang) v0.12.0** - Model Context Protocol (MCP) implementation for Go

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
- **LLama.cpp binaries** (see `/llamacpp/README.MD` for installation)
- **GGUF format models** (see `/models/README.MD` for sources)

## Quick Start

### 1. Clone and Build

```shell script
git clone <repository-url>
cd byte-vision-mcp
go mod tidy
go build -o byte-vision-mcp
```


### 2. Set Up LLama.cpp

Follow the instructions in `/llamacpp/README.MD` to:
- Download prebuilt binaries, or
- Build from source

### 3. Download Models

See `/models/README.MD` for:
- Recommended model sources
- How to download GGUF models
- Model placement instructions

### 4. Configure Environment

Copy the example configuration:
```shell script
cp example-byte-vision-cfg.env byte-vision-cfg.env
```

Edit `byte-vision-cfg.env` to match your setup:
```
# Update paths to match your installation
LLamaCliPath=/path/to/your/llama-cli
ModelFullPathVal=/path/to/your/model.gguf
AppLogPath=/path/to/logs/
```
### 5. Run the Server

```shell script
./byte-vision-mcp
```

The server will start on `http://localhost:8080/mcp-completion` by default.

## Project Structure

```
byte-vision-mcp/
├── llamacpp/          # LLama.cpp binaries and installation guide
├── logs/              # Application and model logs
├── models/            # GGUF model files
├── prompt-cache/      # Cached prompts for performance
├── main.go            # Main MCP server implementation
├── model.go           # Model execution logic
├── types.go           # Configuration structures
├── byte-vision-cfg.env # Your configuration (create from example)
└── example-byte-vision-cfg.env # Example configuration
```

## Configuration

The `byte-vision-cfg.env` file controls all aspects of the server:

### Application Settings
- `AppLogPath`: Directory for log files
- `AppLogFileName`: Log file name
- `HttpPort`: Server port (default `:8080`)
- `EndPoint`: MCP endpoint path (default `/mcp-completion`)
- `TimeOutSeconds`: Request timeout (default `300`)

### LLama.cpp Settings
- `LLamaCliPath`: Path to llama-cli executable
- `ModelFullPathVal`: Path to your GGUF model file
- `CtxSizeVal`: Context window size
- `GPULayersVal`: Number of layers to offload to GPU
- `TemperatureVal`: Generation temperature
- `PredictVal`: Maximum tokens to generate
- And many more LLama.cpp parameters...

## Usage

### MCP Tool: `generate_completion`

The server exposes a single MCP tool that accepts:

**Input:**
```json
{
  "prompt": "Your text prompt here"
}
```


**Output:**
```json
{
  "content": [
    {
      "type": "text",
      "text": "Generated completion text..."
    }
  ]
}
```

### Example with MCP Client

```javascript
// Example MCP client usage
const result = await mcpClient.callTool("generate_completion", {
  prompt: "Write a short story about a robot:"
});
console.log(result.content[0].text);
```

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

See `/logs/README.MD` for log management details.

## Troubleshooting

### Common Issues

1. **"llama-cli not found"**
    - Check `LLamaCliPath` in your `.env` file
    - Ensure the binary has execute permissions

2. **"Model file not found"**
    - Verify `ModelFullPathVal` points to a valid `.gguf` file
    - Check file permissions

3. **Out of memory errors**
    - Reduce `CtxSizeVal`
    - Use a smaller model
    - Increase `GPULayersVal` for GPU offloading

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
```shell script
go mod tidy
go build -o byte-vision-mcp
```


### Running Tests
```shell script
go test ./...
```


### Dependencies
- `github.com/joho/godotenv` - Environment file loading
- `github.com/metoro-io/mcp-golang` - MCP protocol implementation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Add your license information here]

## Support

- Check the individual README files in each subdirectory for specific setup instructions
- Review logs for detailed error information
- Ensure all paths in configuration are absolute and accessible