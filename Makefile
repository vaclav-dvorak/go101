lesson01:
	go run ./cmd/lessons/01

lesson02:
	go run ./cmd/lessons/02

salesman:
	go run ./cmd/salesman

build_lesson01:
	go build -o ./bin/lesson01 ./cmd/lessons/01

build_lesson02:
	go build -o ./bin/lesson02 ./cmd/lessons/02

build_salesman:
	go build -o ./bin/salesman ./cmd/salesman

buildall: build_lesson01 build_lesson02 build_salesman
