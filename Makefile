BINARY_NAME=bin/main

build:
	@go build -o $(BINARY_NAME) ./cmd/api

run: build
	@./$(BINARY_NAME)

clean:
	@rm -f $(BINARY_NAME)

# PHONY targets
.PHONY: build run clean
