export PATH := $(PATH):$(shell go env GOPATH)/bin

# Makefile
.PHONY: run build wire clean

# Default variables
BINARY_NAME=app

wire:
	wire ./cmd

build: wire
	go build -o $(BINARY_NAME) cmd/main.go cmd/application.go cmd/wire_gen.go

run: 
	go run cmd/main.go cmd/application.go cmd/wire_gen.go

clean:
	rm -f $(BINARY_NAME)
	rm -f cmd/wire_gen.go
	go clean

test:
	go test -v ./...

.DEFAULT_GOAL := run