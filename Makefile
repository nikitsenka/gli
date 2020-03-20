GOCMD=go
BINARY_NAME=gli

all: clean lint test build

build:
	go build -o bin/$(BINARY_NAME) ./cmd/gli

clean:
	rm -rf bin

lint:
	golint ./...

test:
	go test ./...