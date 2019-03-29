PROJECT_SCOPE = sl-faas
PROJECT_NAME = step-worker

include .faas

test:
	@ go test ./...

build:
	@ GOOS=linux go build -o ./dist/process ./src/process
