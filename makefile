.PYONY: arm linux build darwin windows run help ui

buildTime ?= $(shell date '+%Y-%m-%d_%H:%M:%S')
buildGoVersion := $(shell go version|awk '{print $$3}')
author := $(shell git config user.name)
LDFLAGS := -X 'main.buildTime=$(buildTime)' -X 'main.buildGoVersion=$(buildGoVersion)' -X 'main.author=$(author)'
SHELL = /bin/bash


build: 
	@CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o promAdmin 

darwin:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o promAdmin 

linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o promAdmin 

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o promAdmin.exe 

run:
	@go run ./main.go

ui:
	@if [ -d "dist" ]; then rm -rf dist; fi && cd web/ui/prom-admin/ && if [ -d "dist" ]; then rm -rf dist; fi \
	&& pnpm install && pnpm run build && mv dist/ ../../../

help:
	@echo "usage: make <option>"
	@echo "    help   : Show help"
	@echo "    build  : Build the binary of this project for current platform"
	@echo "    run	  : run the  project"
	@echo "    linux  : Build the amd64 linux binary of this project"
	@echo "    darwin : Build the arm64 darwin binary of this project"