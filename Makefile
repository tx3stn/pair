BINARY_NAME=pair
DIR ?= ./...
PWD ?= $(shell pwd)
VERSION ?= $(shell head -n 1 VERSION)

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags "-X github.com/tx3stn/pair/cmd.Version=${VERSION}" -o ${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint fmt ${DIR}
	@golangci-lint run --fix ${DIR}

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover

.PHONY: testsum
testsum:
	@CGO_ENABLED=1 gotestsum --format-hide-empty-pkg --format pkgname-and-test-fails -- -race ${DIR}

.PHONY: test-e2e
test-e2e:
	@docker build . -f .docker/bats-tests.Dockerfile -t pair/e2e-tests:local
	@docker run --rm -it -v ${PWD}/.scripts:/code pair/e2e-tests:local bats --verbose-run --formatter pretty /code/e2e-tests
