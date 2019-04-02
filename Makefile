PROJECT_SCOPE = sl-faas
PROJECT_NAME = step-worker

include .faas

install:
	@ dep ensure

build-%:
	@ GOOS=linux go build -ldflags="-s -w" -o ./dist/$* ./src/$*

build: clean
build: build-check
build: build-fan
build: build-process

.PHONY: install build