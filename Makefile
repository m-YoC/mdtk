.DEFAULT_GOAL := help
SHELL=/bin/bash

up: ## Start all container.
	docker compose up -d
down: ## Down all container.
	docker compose down
build: ## Build all container.
	docker compose build
build-plain: ## Build all container (--progress=plain).
	docker compose build --progress=plain

compress-bin: ## compress binary files
	cd sources && source ./mdtk/version.txt && tar -zcvf ../mdtk_bin_v$${VERSION}.tar.gz ./mdtk_bin 
decompress-bin: ## decompress binary files
	@echo 'run command: tar -zxvf ./mdtk_bin_VERSION.tar.gz'

install-linux-amd64: ## Install mdtk to /usr/local/bin/
	cp ./sources/mdtk_bin/linux_amd64/mdtk /usr/local/bin/mdtk
install-linux-arm64: ## Install mdtk to /usr/local/bin/
	cp ./sources/mdtk_bin/linux_arm64/mdtk /usr/local/bin/mdtk

help: ## Display this help screen.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk -F ':.*?## ' '{printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'