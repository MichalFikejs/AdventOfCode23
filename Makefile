.PHONY: build test run

# The name of your application
APP_NAME = aoc2023 

# The go compiler to use
GO = go

build:
		@$(GO) build -o build/$(APP_NAME)

test:
		@$(GO) test -v ./...

run: build
		./build/$(APP_NAME)
