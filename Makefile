BINARY_NAME=cook
GO_BUILD_ENV=CGO_ENABLED=0 GOOOS=linux GOARCH=amd64

all: fmt vet build test

build:
	${GO_BUILD_ENV} go build -v -o ${BINARY_NAME} .

install:
	${GO_BUILD_ENV}

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v

test_coverage:
	go test ./... -coverprofile=coverage.out

clean:
	go clean
	rm ${BINARY_NAME}

