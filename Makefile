CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}

# Including
include .build_info

all: build

build: ## build

build-image: ##buid image

push-image: ## push image

swag-init: ## init swagger
	swag init -g cmd/main.go -o api/docs

run: ## run application
	go run cmd/main.go
