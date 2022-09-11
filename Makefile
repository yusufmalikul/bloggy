BINARY_NAME=bloggy

build:
	go build -o out/${BINARY_NAME} main.go

run:
	out/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -f ${BINARY_NAME}

format:
	gofmt -w *.go cmd/* pkg/*