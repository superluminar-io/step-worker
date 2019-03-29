PROJECT_SCOPE = sl-faas
PROJECT_NAME = step-worker

include .faas

install:
	@ dep ensure

build:
	@ GOOS=linux go build -o ./dist/process ./src/process

.PHONY: install build