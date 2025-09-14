# 🏗️ Structify

**A sleek CLI tool developed in Go to convert JSON structures into Go structs (and more in the future...).**

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

## ✨ Features

- **Instant Conversion**: Transform JSON to Go structs in milliseconds
- **Flexible I/O**: Accept input from stdin or files, output to stdout or files
- **Custom Naming**: Specify your own struct names with the `-name` flag
- **Deterministic Output**: Alphabetically sorted fields for consistent results
- **Cross-Platform**: Pre-built binaries for Windows, Linux, and macOS

## 🛠️ Installation

### Option 1: Download Pre-built Binary (Recommended)

1. Visit the [Releases Page](https://github.com/valreyh/structify/releases)
2. Download the binary for your operating system
3. (Linux/macOS) Make it executable: `chmod +x structify-*`
4. (Optional) Move it to your PATH for global access

### Option 2: Install via Go

```bash
go install github.com/valreyh/structify@latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/valreyh/structify
cd structify
go build -o structify ./cmd/structify
```

## 🚀 Usage

### Basic Examples

**Convert from a file**:

```bash
structify -i input.json -o output.go -name "User"
```

**Convert from stdin**:

```bash
echo '{"name": "John", "age": 30}' | structify -name "Profile"
```

**Pipe from curl**:

```bash
curl https://api.example.com/data.json | structify -name "ApiResponse"
```

## 📋 Example

**Input JSON** (user.json):

```json
{
  "name": "Alice",
  "age": 28,
  "isActive": true,
  "tags": ["admin", "user"]
}
```

**Command**:

```bash
structify -i user.json -o user.go -name "User"
```

**Output** (```user.go```):

```go
package main

type User struct {
  Age      float64 `json:"age"`
  IsActive bool    `json:"isActive"`
  Name     string  `json:"name"`
  Tags     []string `json:"tags"`
}
```

## 🗺️ Roadmap

- [x] **JSON to Go struct** conversion
- [ ] **YAML to Go struct** support
- [ ] **Custom field tags** (`json`, `yaml`, `validate`)
- [ ] **Go struct to JSON** conversion (and many more in the future ...)
- [ ] **Docker image** distribution 🐳
- [ ] **CI/CD pipeline** with GitHub Actions
- [ ] **Package name customization**

⭐ *If you find this tool useful, please give it a star on GitHub! Any suggestions are welcomed !*