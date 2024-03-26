package main

import (
	"os"

	"github.com/tandesko/cards_validator/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
