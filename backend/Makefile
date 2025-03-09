export PATH := $(PATH):$(shell go env GOPATH)/bin

# Makefile
.PHONY: run build wire clean

# Default variables
BINARY_NAME=app

wire:
	wire ./cmd

build: wire
	go build -o $(BINARY_NAME) cmd/main.go cmd/application.go cmd/wire_gen.go

run: wire
	go run cmd/main.go cmd/application.go cmd/wire_gen.go

clean:
	rm -f $(BINARY_NAME)
	rm -f cmd/wire_gen.go
	go clean

test:
	go test -v ./...

generate_mock:
	mockgen -source=internal/developer/repository.go -destination=internal/developer/mocks/developer_repository.go -package=mocks
	mockgen -source=internal/task/repository.go -destination=internal/task/mocks/task_repository.go -package=mocks
	mockgen -source=internal/schedular/repository.go -destination=internal/schedular/mocks/schedular_repository.go -package=mocks

compose-up:
	docker-compose up -d

test-composer-up:
	docker-compose -f docker-compose-test.yml up -d

.DEFAULT_GOAL := run