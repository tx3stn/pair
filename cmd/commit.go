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
			conf, err := config.Get()
			if err != nil {
				return err
			}

			prefix := prompt.NewPrefixSelector(conf.Prefixes, conf.AccessibleMode)

			selected, err := prefix.Select()
			if err != nil {
				return err
			}

			session := pairing.NewSession(pairing.DataDir, time.Now())

			ticketID, err := session.GetTicketID()
			if err != nil {
				return err
			}

			coAuthors, err := session.GetCoAuthors()
			if err != nil {
				return err
			}

			commitMsg := fmt.Sprintf(
				"%s(%s): ",
				selected,
				ticketID,
			)

			msg, err := prompt.EditCommitMessage(commitMsg, coAuthors, conf.AccessibleMode)
			if err != nil {
				return err
			}

			if err := git.Commit(msg, conf.CommitArgs); err != nil {
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
