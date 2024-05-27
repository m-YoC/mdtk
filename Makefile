.DEFAULT_GOAL := help


up: ## Start all container.
	docker compose up -d
down: ## Down all container.
	docker compose down
build: ## Build all container.
	docker compose build
build-plain: ## Build all container (--progress=plain).
	docker compose build --progress=plain

install-linux: ## Install mdtk to /usr/local/bin/
	cp ./sources/compiled/linux_amd64/mdtk /usr/local/bin/mdtk

help: ## Display this help screen.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk -F ':.*?## ' '{printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'