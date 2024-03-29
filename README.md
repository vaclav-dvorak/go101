# go101

![Go version of a Go module](https://img.shields.io/github/go-mod/go-version/vaclav-dvorak/go101.svg)

Attempt to get wisdom from this [book](https://go101.org).

## Useful commands

Everything is handled by Makefile because laziness.

```bash
# go run
make lesson01

#go build
make build_lesson01
./bin/lesson01
```

To build everything

```bash
make buildall
```

## go project structure

```cmd
goproject/
├── bin
├── cmd
├── pkg
└── Makefile
```

The cmd directory will contain the different binaries, separated in directories.

```cmd
cmd/
├── bin1
│   └── main.go
├── bin2
│   └── main.go
└── bin3
    └── main.go
```

pkg directory will contain all your reusable packages. In your case the common code used by the different binaries. This directory can also be named internal, learn more about it here [https://stackoverflow.com/questions/33351387/how-to-use-internal-packages].

```cmd
pkg
├── reusablepackage1
└── reusablepackage2
```

The bin directory is optional, It can be used to store the generated binaries. In case you are generating binaries to $GOBIN this can be omitted.

```cmd
bin/
├── bin1
├── bin2
└── bin3
```

## clivd

Attempt to make poc cli tool. It accepts only one command `version`

```bash
make buildclivd
./bin/clivd version
Version              v0.0.1
Commit               37739b5
Date                 2021-09-30T20:38:24Z
```

## AOC

Solutions to [Advent of Code 2021](https://adventofcode.com/2021) challenges. Can be found under `aoc` directories.

```cmd
src/
├── aoc01
├── ...
└── aocxx
```
