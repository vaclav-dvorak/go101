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
