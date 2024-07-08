SHELL := /bin/bash

.PHONY: help install fmt clean start build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '		help           		Show this help screen.'
	@echo '		start         		Start server.'
	@echo '		build				Build whole project.'
	@echo '		fmt					Format code.'
	@echo '		install				Install tools.'
	@echo '		clean				Clean generated files.'

install:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install mvdan.cc/gofumpt@latest

fmt:
	gofumpt -l -w -extra .

clean:
	rm -rf tmp/* docs/* wire_gen.go

run: build
	./tmp/server.exe

build:  docs/* wire_gen.go ./tmp/server.exe
	go build -o tmp/server.exe

docs/*: delivery/handler/*.go
	swag init

wire_gen.go: wire.go
	wire ./...