BINARY_NAME=jobscraper
.PHONY: test clean build run

test:
	gotest -v ./...

clean:
	rm -f bin/$(BINARY_NAME)

build: clean
	go build -v -o bin/$(BINARY_NAME) ./cmd/main.go

run:
	go run cmd/main.go
