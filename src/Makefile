CGO_ENABLED=1
CC=x86_64-w64-mingw32-cc
OUTPUT_WINDOWS=Installer-Windows.exe
OUTPUT_LINUX=Installer-Linux
THEME=dark

all: help

windows:
	@echo "Building a project for Windows..."
	@CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o $(OUTPUT_WINDOWS) ./windows/*.go

linux:
	@echo "Building a project for  Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(OUTPUT_LINUX) ./linux/*.go

clean:
	@echo "Cleaning..."
	@rm -f $(OUTPUT_WINDOWS) $(OUTPUT_LINUX)

help:
	@echo "Use 'make windows' to build the Windows version."
	@echo "Use 'make linux' to build the Linux version."
	@echo "Use 'make clean' to clean."

.PHONY: all windows linux clean help
