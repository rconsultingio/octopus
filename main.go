package main

import (
	"fmt"
	"os"

	"github.com/rconsultingio/octopus/pkg/cmd"
)

func main() {
	command := cmd.NewDefaultCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
