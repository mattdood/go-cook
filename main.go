package main

import (
	"os"

	"github.com/mattdood/go-cook/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
