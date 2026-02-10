package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdCommit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			conf, err := config.Get()
			if err != nil {
				return err
			}

			prefix := prompt.NewPrefixSelector(conf.Prefixes, conf.AccessibleMode)

			commitType, err := prefix.Select()
			if err != nil {
				return err
			}

			session := pairing.NewSession(pairing.DataDir, time.Now())

			ticketID, err := session.GetTicketID()
			if err != nil {
				return err
			}

			if ticketID == "" {
				ticketID, err = setTicketID(session, conf, "")
				if err != nil {
					return err
				}
			}

			coAuthors, err := session.GetCoAuthors()
			if err != nil {
				return err
			}

			if len(coAuthors) == 0 {
				coAuthors, err = setCoAuthors(session, conf)
				if err != nil {
					return err
				}
			}

			msg, err := prompt.EditCommitMessage(
				fmt.Sprintf("%s(%s%s): ", commitType, conf.TicketPrefix, ticketID),
				coAuthors,
				conf.AccessibleMode,
			)
			if err != nil {
				return err
			}

			if _, err := git.Commit(ctx, msg, conf.CommitArgs); err != nil {
				return err
			}

			log.Printf("committed:\n%s", msg)

			return nil
		},
		Short: "Commit your work",
		Use:   "commit",
	}

	return cmd
}
