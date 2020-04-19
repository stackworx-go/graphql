 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: fmt build lint test
build:
	$(GOBUILD) -v cmd/main.go
test:
	$(GOTEST) -short -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
lint:
	golangci-lint run ./...
fmt:
	go fmt ./...
