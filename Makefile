.DEFAULT_GOAL := help

SHELL := /bin/bash

## help: Display this help message
help: Makefile
	@sed -n 's/^##//p' $<
