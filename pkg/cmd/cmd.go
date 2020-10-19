package cmd

import (
	"io"
	"os"

	"github.com/rconsultingio/octopus/pkg/cmd/env"
	"github.com/spf13/cobra"
)

func NewDefaultCommand() *cobra.Command {
	return newCommand(os.Stdin, os.Stdout, os.Stderr)
}

func newCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "octopus",
		Short: "octopus is a simple ci/cd cli",
		Long:  "octopus is a simple ci/cd cli",
	}

	cmd.AddCommand(env.NewCmdEnv())

	return cmd
}
