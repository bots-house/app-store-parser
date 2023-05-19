package main

import (
	"fmt"
	"os"

	"bots-house/asp/cli"
)

func main() {
	if err := cli.New().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
