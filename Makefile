CGO_ENABLED=1
CC=x86_64-w64-mingw32-cc
OUTPUT_WINDOWS=./bin/Installer-Windows.exe
OUTPUT_LINUX=./bin/Installer-Linux
THEME=dark

windows:
	@echo "Building a project for Windows..."
	@CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o $(OUTPUT_WINDOWS) ./src/*.go

linux:
	@echo "Building a project for Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_LINUX) ./src/*.go

all:
	@echo "Building a project for Windows..."
	@CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o $(OUTPUT_WINDOWS) ./src/*.go
	@echo "Building a project for Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_LINUX) ./src/*.go

clean:
	@echo "Cleaning..."
	@rm -f $(OUTPUT_WINDOWS) $(OUTPUT_LINUX)

help:
	@echo "Use 'make windows' to build the Windows version."
	@echo "Use 'make linux' to build the Linux version."
	@echo "Use 'make clean' to clean."

.PHONY: all windows linux clean help
