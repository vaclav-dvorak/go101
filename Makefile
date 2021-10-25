GIT_REV    ?= $(shell git rev-parse --short HEAD)
DATE       ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION    ?= v0.0.1

lesson01:
	go run ./src/lesson01

lesson02:
	go run ./src/lesson02

salesman:
	go run ./src/salesman

basket:
	go run ./src/basket

build_lesson01:
	go build -ldflags="-s -w" -o ./bin/lesson01 ./src/lesson01

build_lesson02:
	go build -ldflags="-s -w" -o ./bin/lesson02 ./src/lesson02

build_salesman:
	go build -ldflags="-s -w" -o ./bin/salesman ./src/salesman

buildall: build_lesson01 build_lesson02 build_salesman

buildclivd:  ## Builds the CLI
	@go build \
	-ldflags "-s -w -X github.com/vaclav-dvorak/go101/src/clivd/cmd.version=${VERSION} -X github.com/vaclav-dvorak/go101/src/clivd/cmd.commit=${GIT_REV} -X github.com/vaclav-dvorak/go101/src/clivd/cmd.date=${DATE}" \
	-a -o ./bin/clivd ./src/clivd
