SHELL := bash

PROJECT_NAME  ?= ezuzu
GO_BINARY_DIR := bin
CGO_ENABLED   := 0
GOOS          ?= linux
GOARCH        ?= arm64
MIGRATIONS_DIR = business/data/migrations/services
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

.PHONY    : migration
migration :
	@cd $(MIGRATIONS_DIR)/${SERVICE} && goose postgres "host=localhost user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} port=${DB_PORT}" up

.PHONY : test
test   :
	@CGO_ENABLED=$(CGO_ENABLED) go test -v ./... --cover
