## Go
.PHONY: setup test build run
setup:
	@echo "Start setup"
	@env go get ./...

build: setup
	@echo "build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/api ./cmd

run: setup
	@echo "go run"
	@go run ./cmd/main.go
