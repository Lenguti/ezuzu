SHELL := bash

PROJECT_NAME  ?= ezuzu
GO_BINARY_DIR := bin
CGO_ENABLED   := 0
GOOS          ?= linux
GOARCH        ?= arm64
MIGRATIONS_DIR = business/data/migrations
SERVICES_DIR   = app/services

.PHONY : build
build  :
	@echo foo

.PHONY : run
run    :
	@docker compose up -d --build

.PHONY : stop
stop   :
	@docker compose down

.PHONY : test
test   :
	@CGO_ENABLED=$(CGO_ENABLED) go test ./... --cover
