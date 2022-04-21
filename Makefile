.PHONY: fmt test tests vet

fmt:
	go fmt ./...

test:
	go test -race -covermode=atomic ./...

tests: fmt vet test

vet:
	go vet ./...
