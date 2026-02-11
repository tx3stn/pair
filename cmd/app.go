// Package cmd contains all of the CLI commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/flags"
	"github.com/tx3stn/pair/internal/logger"
)

// Version is the CLI version set via linker flags at build time.
var Version string

func NewApp() *cobra.Command {
	cfg, err := config.Get()
	if err != nil {
		cfg = &config.Config{}
	}

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

	rootCmd.AddCommand(NewCmdCommit(cfg))
	rootCmd.AddCommand(NewCmdDone(cfg))
	rootCmd.AddCommand(NewCmdOn(cfg))
	rootCmd.AddCommand(NewCmdWith(cfg))

	rootCmd.PersistentFlags().
		BoolVar(&flags.Verbose, "verbose", false, "display verbose output for more detail on what the command is doing")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		logger.New(cfg.AccessibleMode)
	}

	return rootCmd
}
