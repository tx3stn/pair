package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
)

func NewCmdCur(_ *config.Config) *cobra.Command {
	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir)

			ticket, coAuthors, err := session.Current()
			if err != nil {
				return err
			}

			numCoAuthors := len(coAuthors)
			if ticket == "" && numCoAuthors == 0 {
				slog.Info("no active pairing session")

				return nil
			}

			slog.Info("pairing on", "ticketID", ticket)
			slog.Info(fmt.Sprintf("with %d coauthors", numCoAuthors), "coAuthors", coAuthors)

			return nil
		},
		Short: "return the current active session details (cur as in current)",
		Use:   "cur",
	}
}
