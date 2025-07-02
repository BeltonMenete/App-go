.PHONY: run build clean test

all: run

run:
	@echo "Running Go application..."
	go run .

build:
	@echo "Building Go application..."
	go build -o myapp .
	@echo "Executable 'myapp' created."

clean:
	@echo "Cleaning up..."
	rm -f myapp

test:
	@echo "Running Go tests..."
	go test ./...