package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdNew(conf *config.Config) *cobra.Command {
	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir)

			if err := session.Clean(); err != nil {
				return err
			}

			_, err := setCoAuthors(session, conf, false)
			if !errors.Is(err, prompt.ErrNoCoAuthorsSelected) {
				return err
			}

			_, err = setTicketID(session, conf, "")

			return err
		},
		Short: "start a new session, overwritting current values",
		Use:   "new",
	}
}
