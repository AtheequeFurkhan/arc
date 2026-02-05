package main

import (
	"os"

	"github.com/atheeque/arc/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
