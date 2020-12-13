build:
	@docker build -t mollie-cli:latest -f Dockerfile .
.PHONY: build

run:
	@docker run --rm mollie-cli:latest
.PHONY: run

lint:
	@go version

	@echo "Running go lint"
	@golint ./...

	@echo "Running go vet"
	@go vet ./...
.PHONY: lint

test: run
.PHONY: test

test-local:
	@go test -v ./... -coverprofile cover.out
.PHONY: test-local

coverage:
	@go test ./... -coverprofile cover.out
	@go tool cover -func=cover.out
.PHONY:  coverage

clean:
	@go mod verify
	@go mod tidy
.PHONY: clean

compile-master:
	@go build -o mollie-master ./cmd/mollie/main.go
.PHONY: compile-master

compile-current:
	@go build -o mollie ./cmd/mollie/main.go
.PHONY: compile-current