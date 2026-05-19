package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
)

func NewCmdNew(conf *config.Config) *cobra.Command {
	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir)

			if err := session.Clean(); err != nil {
				return err
			}

			if _, err := setCoAuthors(session, conf, false); err != nil {
				return err
			}

			_, err := setTicketID(session, conf, "")

			return err
		},
		Short: "start a new session, overwritting current values",
		Use:   "new",
	}
}
