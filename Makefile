BUILD_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo no-commit-id)
IMAGE_NAME := anubhavmishra/hello-oscon

.DEFAULT_GOAL := help
help: ## Show available targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean the project
	rm -rf ./build
	mkdir ./build

deps:
	go get .

build-service: ## Build the project
	mkdir -p ./build/linux/amd64
	GOOS=linux GOARCH=amd64 go build -v -o ./build/linux/amd64/hello-oscon .
	#docker build -t $(IMAGE_NAME):$(BUILD_ID) .
	#docker tag $(IMAGE_NAME):$(BUILD_ID) $(IMAGE_NAME):latest

push: ## Docker push the service images tagged 'latest' & 'BUILD_ID'
	docker push $(IMAGE_NAME):$(BUILD_ID)
	docker push $(IMAGE_NAME):latest

deps-test:
	go get -t

test: ## Run tests, coverage reports, and clean (coverage taints the compiled code)
	go test -v .

run: ## Build and run the project
	mkdir -p ./build
	go build -o ./build/hello-oscon && ./build/hello-oscon