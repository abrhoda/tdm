PACKAGE ?= ./...
CLI_PACKAGE ?= ./cmd/cli/main.go
EXECUTABLE_NAME ?= tdm
OUT_DIR ?= ./out

.PHONY: build-pf2e-all
build-pf2e-all: ## builds the pf2e dataset without restrictions
	@mkdir -p $(OUT_DIR)
	@go build -o $(OUT_DIR)/$(EXECUTABLE_NAME) $(PACKAGE)

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


.PHONY: help
help: ## print this help message
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "\033[36m%-20s\033[0m %s\n", $$1, $$NF}' $(MAKEFILE_LIST)
