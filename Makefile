SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: build
build: clean backend/bin/server

.PHONY: clean
clean:
	@rm -f backend/bin/*

backend/bin/server: backend/cmd/bingo-local/main.go
	@cd backend
	@go build -o bin/server ./cmd/bingo-local
