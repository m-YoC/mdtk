.DEFAULT_GOAL := help

hello: ## golang hello world test
	cd hello-world && go run hello.go
run: ## mdtk run
	cd mdtk && go run mdtk.go test -f ../Taskfile.md -- arg1="aaaa a" arg2=bbbb
runhelp: ## mdtk run task help
	cd mdtk && go run mdtk.go help -f ../Taskfile.md -- arg1="aaaa a" arg2=bbbb
build: ## build mdtk (test)
	cd mdtk && CGO_ENABLED=0 go build -o ../mdtk -ldflags '-s -w' -trimpath
install: ## mdtk install (test)
	cd mdtk && CGO_ENABLED=0 go install -ldflags '-s -w' -trimpath mdtk.go
test:
	@cd mdtk && go test ./...

	

cross-compile: ## cross compile
	@mkdir -p ./mdtk_bin/linux_amd64 && cd mdtk && GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o ../mdtk_bin/linux_amd64/mdtk -ldflags '-s -w' -trimpath
	@mkdir -p ./mdtk_bin/linux_arm64 && cd mdtk && GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -o ../mdtk_bin/linux_arm64/mdtk -ldflags '-s -w' -trimpath	
	@mkdir -p ./mdtk_bin/win_amd64   && cd mdtk && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ../mdtk_bin/win_amd64/mdtk.exe -ldflags '-s -w' -trimpath

# @mkdir -p ./mdtk_bin/mac_amd64   && cd mdtk && GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o ../mdtk_bin/mac_amd64/mdtk -ldflags '-s -w' -trimpath
# @mkdir -p ./mdtk_bin/mac_amd64   && cd mdtk && GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -o ../mdtk_bin/mac_arm64/mdtk -ldflags '-s -w' -trimpath

install-go-licenses:
	@GO111MODULE=on go install github.com/google/go-licenses@latest
	@GO111MODULE=on go install github.com/Songmu/gocredits/cmd/gocredits@latest
check-licenses:
	@cd mdtk && go-licenses csv .
write-licenses:
	@cd mdtk && gocredits > ../CREDITS

help: ## Display this help screen.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk -F ':.*?## ' '{printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
