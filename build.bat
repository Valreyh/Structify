@echo off
echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o bin/structify-windows-amd64.exe ./cmd/structify

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o bin/structify-linux-amd64 ./cmd/structify

echo Building for macOS (Intel)...
set GOOS=darwin
set GOARCH=amd64
go build -o bin/structify-darwin-amd64 ./cmd/structify

echo Building for macOS (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build -o bin/structify-darwin-arm64 ./cmd/structify

echo.
echo Build complete! Check the bin/ folder.
echo.
pause