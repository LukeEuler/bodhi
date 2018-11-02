all: build

GLIDE := $(shell command -v glide 2> /dev/null)
GOFMT := $(shell command -v gofmt 2> /dev/null)
GOIMPORTS := $(shell command -v goimports 2> /dev/null)
GOLINT :=  $(shell command -v golint 2> /dev/null)
CILINT := $(shell command -v golangci-lint 2> /dev/null)

style: clean
ifndef GOFMT
	$(error "gofmt is not available please install gofmt")
endif
ifndef GOIMPORTS
	$(error "goimports is not available please install goimports")
endif
	@echo ">> checking code style"
	@! find . -path ./vendor -prune -o -name '*.go' -print | xargs gofmt -d | grep '^'
	@! find . -path ./vendor -prune -o -name '*.go' -print | xargs goimports -d | grep '^'

format:
ifndef GOFMT
	$(error "gofmt is not available please install gofmt")
endif
ifndef GOIMPORTS
	$(error "goimports is not available please install goimports")
endif
	@echo ">> formatting code"
	@glide nv | xargs go fmt
	@find . -path ./vendor -prune -o -name '*.go' -print | xargs goimports -l | xargs goimports -w
	@echo ">> done"

lint:
ifndef GOLINT
	$(error "golint is not available please install golint")
endif
	@echo ">> checking code lint"
	@! go list ./... | grep -v -e "updater\/plugin\/trx\/api" -e "updater\/plugin\/trx\/core" | xargs golint | sed "s:^$(CURRENT_DIR)/::" | grep '^'

cilint:
ifndef CILINT
	$(error "golangci-lint is not available please install golangci-lint")
endif
	@golangci-lint run

test: style lint
	@echo ">> testing"
	@glide nv | xargs go test -cover

build: test
	go build -o build/bodhi example/main.go

linux: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bodhi example/main.go

clean:
	@rm -rf build