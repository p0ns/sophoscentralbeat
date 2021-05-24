package main

import (
	"os"

	"github.com/p0ns/sophoscentralbeat/cmd"

	_ "github.com/p0ns/sophoscentralbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
