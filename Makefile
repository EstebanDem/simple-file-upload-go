BINARY_NAME=bin/main
UPLOADS_DIR=../uploads

build:
	@go build -o $(BINARY_NAME) ./cmd/api

run: create-uploads-dir build
	@./$(BINARY_NAME)

create-uploads-dir:
	@mkdir -p $(UPLOADS_DIR)

clean:
	@rm -f $(BINARY_NAME)
	@rm -rf $(UPLOADS_DIR)

# PHONY targets
.PHONY: build run clean
