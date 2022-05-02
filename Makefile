.DEFAULT_GOAL := help

SHELL := /bin/bash

APP_VERSION := 0.0.1
GIT_COMMIT := $(shell git rev-parse --short HEAD)

IMAGE_NAME := ventilo/elliot-api
IMAGE_TAG := $(APP_VERSION)-$(GIT_COMMIT)

## help: Display this help message
help: Makefile
	@sed -n 's/^##//p' $<

build:
	GOOS=linux GOARCH=amd64 go build -o .build/app
	docker build --tag $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
