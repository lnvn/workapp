.PHONY: build

build:
	@echo "Building binary..."
	go build -o ./main ./main.go

.PHONY: build-docker

build-docker:
	@echo "Building docker image..."
	docker build -t test:latest .