package cmd

import (
	"log/slog"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
)

func NewCmdDone(_ *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir, time.Now())

			if err := session.Clean(); err != nil {
				return err
			}

			slog.Info("pairing session complete!")

			return nil
		},
		Short: "All done - remove the active pairing state",
		Use:   "done",
	}

	return cmd
}
