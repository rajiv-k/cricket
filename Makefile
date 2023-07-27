GOOS ?= linux
CGO_ENABLED ?= 0
BINARY_NAME ?= cricket
BUILD_DIR ?= ./build

YELLOW=\033[33m
RESET=\033[0m

cricket: ## Build cricket binary
	@mkdir -p ${BUILD_DIR}
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} go build -o ${BUILD_DIR}/${BINARY_NAME} ./cmd/cricket

clean: ## Clean up build artifact(s) and remove ${BUILD_DIR}
	rm -rf build

help: show-defaults## Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

show-defaults:
	@echo "-----------------------------------------------------------------------------"
	@printf "BUILD_DIR: ${YELLOW}${BUILD_DIR}${RESET}\n"
	@printf "GOOS: ${YELLOW}${GOOS}${RESET}\n"
	@printf "BINARY_NAME: ${YELLOW}${BINARY_NAME}${RESET}\n"
	@echo "-----------------------------------------------------------------------------"
	@echo

.PHONY: show-defaults clean help
