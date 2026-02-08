// Package cmd contains all of the CLI commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/flags"
)

// Version is the CLI version set via linker flags at build time.
var Version string

func NewApp() *cobra.Command {
	rootCmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			err := ccmd.Help()
			if err != nil {
				return fmt.Errorf("error getting cobra help: %w", err)
			}

			return nil
		},
		Short:   "Your simple pair commit helper.",
		Use:     "pair",
		Version: Version,
	}

	rootCmd.AddCommand(NewCmdOn())

	rootCmd.PersistentFlags().
		BoolVar(&flags.Verbose, "verbose", false, "display verbose output for more detail on what the command is doing")

	return rootCmd
}
