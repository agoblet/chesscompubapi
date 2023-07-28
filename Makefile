all: test lint cov

test:
	go test -v -race -coverprofile=coverage.out -covermode=atomic -tags e2e

lint:
	golangci-lint -v run

cov:
	go tool cover -html=coverage.out
