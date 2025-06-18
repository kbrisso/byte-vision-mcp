# LlamaCpp Installation Guide
This MCP server requires LlamaCpp binaries to function. You can either download prebuilt binaries or build from source.
## Option 1: Download Prebuilt Binaries (Recommended)
### Windows
1. Visit the [LlamaCpp releases page](https://github.com/ggerganov/llama.cpp/releases)
2. Download the latest release for Windows:
    - For CUDA GPU support: `llama-*-bin-win-cuda-cu*.zip`
    - For CPU only: `llama-*-bin-win-x64.zip`
    - For AMD GPU (ROCm): `llama-*-bin-win-rocm.zip`

3. Extract the ZIP file to your desired location (e.g., `byte-vision-mcp/llamacpp/`)
4. The main executable will be `llama-cli.exe`

### Linux
1. Visit the [LlamaCpp releases page](https://github.com/ggerganov/llama.cpp/releases)
2. Download the appropriate release:
    - For CUDA GPU support: `llama-*-bin-ubuntu-x64-cuda-cu*.tar.gz`
    - For CPU only: `llama-*-bin-ubuntu-x64.tar.gz`

3. Extract: `tar -xzf llama-*-bin-ubuntu-x64.tar.gz`
4. The main executable will be `llama-cli`

### macOS
1. Install via Homebrew (recommended):
``` bash
   brew install llama.cpp
```
1. Or download from releases page and extract manually

## Option 2: Build from Source
### Prerequisites
- **Windows**: Visual Studio 2019+ with C++ support, or MinGW-w64
- **Linux**: GCC 8+ or Clang 6+
- **macOS**: Xcode command line tools
- **Git** for cloning the repository
- **CMake** 3.14 or higher

### Build Steps
1. **Clone the repository:**
``` bash
   git clone https://github.com/ggerganov/llama.cpp.git
   cd llama.cpp
```
1. **Build with CMake:**
   **For CPU only:**
``` bash
   mkdir build
   cd build
   cmake ..
   cmake --build . --config Release
```
**For NVIDIA GPU (CUDA) support:**
``` bash
   mkdir build
   cd build
   cmake .. -DLLAMA_CUDA=ON
   cmake --build . --config Release
```
**For AMD GPU (ROCm) support:**
``` bash
   mkdir build
   cd build
   cmake .. -DLLAMA_HIPBLAS=ON
   cmake --build . --config Release
```
**For Metal (macOS) support:**
``` bash
   mkdir build
   cd build
   cmake .. -DLLAMA_METAL=ON
   cmake --build . --config Release
```
1. **Locate the built executable:**
    - **Windows**: `build/bin/Release/llama-cli.exe`
    - **Linux/macOS**: `build/bin/llama-cli`

## Configuration
After installing LlamaCpp, update your file with the correct path: `byte-vision-cfg.env`
``` env
# Update this path to point to your llama-cli executable
LLamaCliPath=/byte-vision-mcp/llamacpp/llama-cli.exe
```
**Examples:**
- Windows: `/byte-vision-mcp/llamacpp/llama-cli.exe`
- Linux: `/home/user/llamacpp/llama-cli`
- macOS: `/usr/local/bin/llama-cli` (if installed via Homebrew)
