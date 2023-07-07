all: build test lint

build:
	go build -v

test:
	go test -v

lint:
	golangci-lint -v run
