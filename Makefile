BINARY_NAME=pair
DIR ?= ./...
PWD ?= $(shell pwd)
VERSION ?= $(shell head -n 1 VERSION)

define ajv-docker
	docker run --rm -v "${PWD}":/repo weibeld/ajv-cli:5.0.0 ajv --spec draft2020
endef

define vhs-docker
	docker run --rm -v "${PWD}":/vhs --workdir /vhs ${BINARY_NAME}/vhs:local
endef

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags "-s -w -X github.com/tx3stn/pair/cmd.Version=${VERSION}" -o ${BINARY_NAME}

.PHONY: generate-gifs
generate-gifs: build
	@docker build --tag ${BINARY_NAME}/vhs:local -f ./.docker/demo-gif.Dockerfile .
	@$(vhs-docker) /vhs/.scripts/gifs/overview.tape
	@$(vhs-docker) /vhs/.scripts/gifs/commit.tape
	@$(vhs-docker) /vhs/.scripts/gifs/done.tape
	@$(vhs-docker) /vhs/.scripts/gifs/on.tape
	@$(vhs-docker) /vhs/.scripts/gifs/with.tape

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint fmt ${DIR}
	@golangci-lint run --fix ${DIR}

.PHONY: schema-example-lint
schema-example-lint:
	@$(ajv-docker) validate -s /repo/.schema/schema.json -d /repo/.schema/pair.json

.PHONY: schema-validate
schema-validate:
	@$(ajv-docker) compile -s /repo/.schema/schema.json

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover

.PHONY: testsum
testsum:
	@CGO_ENABLED=1 gotestsum --format-hide-empty-pkg --format pkgname-and-test-fails -- -race ${DIR}

.PHONY: test-e2e
test-e2e:
	@docker build . -f .docker/bats-tests.Dockerfile -t pair/e2e-tests:local
	@docker run --rm -e TERM=xterm -v ${PWD}/.scripts:/code pair/e2e-tests:local bats --verbose-run --formatter pretty /code/e2e-tests
