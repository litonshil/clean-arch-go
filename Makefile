PROJECT_NAME := smart_mess
PKG_LIST := $(shell go list ${PROJECT_NAME}/... | grep -v /vendor/)

.PHONY: all dep build clean test help development

all: build ## Build the project

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

########################
### DEVELOP and TEST ###
########################

development: ## Set up development environment
	# Booting up dependency containers
	@docker-compose up -d consul db redis

	# Wait for consul container to be ready
	@while ! curl --request GET -sL --url 'http://localhost:8500/' > /dev/null 2>&1; do printf .; sleep 1; done

	# Setting KV, dependency of app
	@curl --request PUT --data-binary @config.local.json http://localhost:8500/v1/kv/${PROJECT_NAME} || { echo "Failed to set KV"; exit 1; }

	# Building clean-arch
	@docker-compose up --build ${PROJECT_NAME}

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)
	@docker-compose down

build: ## Build the project
	# Build commands here if needed
	@echo "Build completed"
