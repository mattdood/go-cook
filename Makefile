BINARY_NAME=cook
GO_BUILD_ENV=CGO_ENABLED=0 GOOOS=linux GOARCH=amd64

all: fmt vet lint test build  ## fmt, vet, lint, test, build

build:  ## Build binary package
	${GO_BUILD_ENV} go build -v -o ${BINARY_NAME} .

install:  ## Install package with go install
	${GO_BUILD_ENV} go install

fmt:  ## Run gofmt on all files
	go fmt ./...

vet:  ## Run govet on all files
	go vet ./...

lint:  ## Run linter on all files
	golangci-lint run

deps: lint-dep  ## Dependency installs

lint-dep:  ## Golangci-lint dependency install
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
	golangci-lint --version

test: ## Run and output test results for all files
	go test ./... -v

test_coverage:  ## Output test coverage
	go test ./... -coverprofile=coverage.out

clean:  ## Go clean + binary removal
	go clean
	rm ${BINARY_NAME}

help:  ## Shows available Makefile commands in a list
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
