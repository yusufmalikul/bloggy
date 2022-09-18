BINARY_NAME=bloggy

PACKAGE=github.com/yusufmalikul/bloggy
VERSION=1.0.0
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP=$(shell date -u)
LDFLAGS="-X '${PACKAGE}/cmd.Version=${VERSION}' -X '${PACKAGE}/cmd.Commit=${COMMIT}' -X '${PACKAGE}/cmd.BuildTimestamp=${BUILD_TIMESTAMP}' -s -w"

build:
	go build -ldflags=${LDFLAGS} -o out/${BINARY_NAME} main.go

run:
	out/${BINARY_NAME}

build_and_run: build run

release:
	GOOS=darwin go build -trimpath -ldflags=${LDFLAGS} -o out/darwin/${BINARY_NAME} main.go
	GOOS=linux go build -trimpath -ldflags=${LDFLAGS} -o out/linux/${BINARY_NAME} main.go
	GOOS=windows go build -trimpath -ldflags=${LDFLAGS} -o out/windows/${BINARY_NAME}.exe main.go

clean:
	go clean
	rm -f ${BINARY_NAME}

format:
	gofmt -w *.go cmd/* pkg/*
