BINARY := kubectl-ns

all: build

build:
	CGO_ENABLED=0 go build -o $(BINARY) .
clean:
	rm -f $(BINARY)
install: build
	sudo mv $(BINARY) /usr/local/bin/
help:
	@echo "Makefile commands:"
	@echo "  make build        Build the static binary"
	@echo "  make clean        Remove the built binary"
	@echo "  make install      Install the binary to /usr/local/bin"
	@echo "  make help         Show this help message"
	@echo "  make ARCH=arm64   Build for ARM64 architecture (64-bit)"
	@echo "  make ARCH=amd64   Build for AMD64 architecture (default)"

.PHONY: all build clean install help
