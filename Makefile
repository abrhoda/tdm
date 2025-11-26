PACKAGE ?= ./...
CLI_PACKAGE ?= ./cmd/cli/main.go
EXECUTABLE_NAME ?= tdm
OUT_DIR ?= ./out
PODMAN_DIR ?= ./podman

.PHONY: build-cli
build-cli: ## builds the cli binary 
	@mkdir -p $(OUT_DIR)
	@go build -o $(OUT_DIR)/$(EXECUTABLE_NAME) $(CLI_PACKAGE)

.PHONY: test
test: ## run all test 
	@go test $(PACKAGE) -v

.PHONY: format
format: ## format project
	@go fmt $(PACKAGE)


.PHONY: vet
vet: ## runs a `go vet` check for the project
	@go vet

.PHONY: clean
clean: ## remove `./out/` and all files in it
	@rm -rf $(OUT_DIR)

.PHONY: start-db
start-db: ## starts the docker container for postgres and pgadmin using the compose file defined in the `./docker/ dir
	@podman compose -f $(PODMAN_DIR)/local-compose.yaml up --detach --build

.PHONY: stop-db
stop-db: ## stops the docker containers for postgres and pgadmin
	@podman compose -f $(PODMAN_DIR)/local-compose.yaml down --volumes

.PHONY: restart-podman-socket
restart-podman-socket: ## podman socket giving troubles? jump start it.
	@systemctl --user stop podman.socket podman.service
	@systemctl --user start podman.socket podman.service

.PHONY: help
help: ## print this help message
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "\033[36m%-20s\033[0m %s\n", $$1, $$NF}' $(MAKEFILE_LIST)
