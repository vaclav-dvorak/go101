package main

import (
	"log"

	"github.com/vaclav-dvorak/go101/src/clivd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Print(err.Error())
	}
}
