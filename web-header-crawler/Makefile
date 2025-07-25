IMAGE_NAME := crawler
IMAGE_TAG := latest

.PHONY: build

build:
	@echo "Building binary..."
	go build -o ./main ./main.go

.PHONY: build-docker

build-docker:
	@echo "Building docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

.PHONY: test-docker

test-docker:
	@echo "Running test docker image..."
	docker run --rm $(IMAGE_NAME):$(IMAGE_TAG)