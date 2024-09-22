# Makefile for GO AORA API

# Install
install:
	@go mod tidy

# Run the application
run:
	@go run cmd/api/main.go