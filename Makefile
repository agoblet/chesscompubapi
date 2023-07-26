all: test lint cov

test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic

lint:
	golangci-lint -v run

cov:
	go tool cover -html=coverage.out
