GIT_REV    ?= $(shell git rev-parse --short HEAD)
DATE       ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION    ?= v0.0.1

lesson01:
	go run ./cmd/lessons/01

lesson02:
	go run ./cmd/lessons/02

salesman:
	go run ./cmd/salesman

basket:
	go run ./cmd/basket

build_lesson01:
	go build -o ./bin/lesson01 ./cmd/lessons/01

build_lesson02:
	go build -o ./bin/lesson02 ./cmd/lessons/02

build_salesman:
	go build -o ./bin/salesman ./cmd/salesman

buildall: build_lesson01 build_lesson02 build_salesman

buildclivd:  ## Builds the CLI
	@go build \
	-ldflags "-X github.com/vaclav-dvorak/go101/cmd/clivd/cmd.version=${VERSION} -X github.com/vaclav-dvorak/go101/cmd/clivd/cmd.commit=${GIT_REV} -X github.com/vaclav-dvorak/go101/cmd/clivd/cmd.date=${DATE}" \
	-a -o ./bin/clivd ./cmd/clivd
