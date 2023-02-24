package main

import (
	"os"

	"github.com/svrakitin/gossip-glomers/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
