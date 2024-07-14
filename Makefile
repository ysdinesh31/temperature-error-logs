# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=temperature-error-logs
BINARY_PATH=./bin/$(BINARY_NAME)

.PHONY: all
all: build

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_PATH) -v ./cmd/temperature-error-logs

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: run
run: build
	$(BINARY_PATH)




# Help
.PHONY: help
help:
	@echo "Choose a command run in "$(BINARY_NAME)":"
	@echo "  make all       - builds the project"
	@echo "  make build     - Builds the project"
	@echo "  make clean     - Cleans the project"
	@echo "  make run       - Runs the project"
	@echo "  make help      - Displays this help message"
