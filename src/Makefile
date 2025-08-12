# Makefile

# Default Target 
all: build


# Compile
build:
	# x86_64
	GOOS=linux GOARCH=amd64 go build -o go-digestive main.go

build_mac_arm:
	# macOS ARM64 (native M1/M2).
	GOOS=darwin GOARCH=arm64 go build -o go-digestive main.go
	

.PHONY: all build 