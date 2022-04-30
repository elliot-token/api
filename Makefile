.DEFAULT_GOAL := help

SHELL := /bin/bash

IMAGE_NAME := ventilo/elliot-api
APP_VERSION := 0.0.1

## help: Display this help message
help: Makefile
	@sed -n 's/^##//p' $<

build:
	GOOS=linux GOARCH=amd64 go build -o build/app
	docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) --tag $(IMAGE_NAME):$(APP_VERSION) .

push:
	docker push