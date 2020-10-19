package env

import (
	"fmt"
	"os"

	"github.com/rconsultingio/octopus/pkg/manifest"
	"github.com/spf13/cobra"
)

const (
	defaultManifest = "octopus.yml"
)

type EnvOptions struct {
	Filename string
	Export   bool
}

func NewCmdEnv() *cobra.Command {
	o := &EnvOptions{
		Filename: defaultManifest,
	}

	cmd := &cobra.Command{
		Use: "env",
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.run(); err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&o.Filename, "filename", "f", o.Filename, "The manifest filename")
	cmd.Flags().BoolVar(&o.Export, "export", o.Export, "Exports environment variables in the manifest")

	cmd.MarkFlagRequired("filename")

	return cmd
}

func (o *EnvOptions) run() error {
	m, err := manifest.LoadFromFile(o.Filename)
	if err != nil {
		return err
	}

	if o.Export {
		if err := export(m); err != nil {
			return fmt.Errorf("can't export environment variables, %w", err)
		}
		return nil
	}

	return nil
}

func export(m *manifest.Manifest) error {
	for name, value := range m.Env {
		fmt.Printf("%s=%s\n", name, value)
	}

	return nil
}
