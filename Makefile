all: build

MODULE = "github.com/LukeEuler/bodhi"

GOIMPORTS := $(shell command -v goimports 2> /dev/null)
CILINT := $(shell command -v golangci-lint 2> /dev/null)

style: clean
ifndef GOIMPORTS
	$(error "goimports is not available please install goimports")
endif
	! find . -path ./vendor -prune -o -name '*.go' -print | xargs goimports -d -local ${MODULE} | grep '^'

format:
ifndef GOIMPORTS
	$(error "goimports is not available please install goimports")
endif
	find . -path ./vendor -prune -o -name '*.go' -print | xargs goimports -l -local ${MODULE} | xargs goimports -l -local ${MODULE} -w

cilint:
ifndef CILINT
	$(error "golangci-lint is not available please install golangci-lint")
endif
	golangci-lint run

test: style cilint
	go test -cover ./...

build: test
	go build -o build/bodhi example/main.go

linux: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bodhi example/main.go

clean:
	@rm -rf build

.PHONY: all style format cilint test build linux