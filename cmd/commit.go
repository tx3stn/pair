package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdCommit(conf *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

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
				err := setCoAuthors(session, conf)
				if err != nil && !errors.Is(err, prompt.ErrNoCoAuthorsSelected) {
					return err
				}
			}

			if !strings.HasPrefix(ticketID, conf.TicketPrefix) {
				ticketID = conf.TicketPrefix + ticketID
			}

			msg, err := prompt.EditCommitMessage(
				fmt.Sprintf("%s(%s): ", commitType, ticketID),
				coAuthors,
				conf.AccessibleMode,
			)
			if err != nil {
				return err
			}

			if _, err := git.Commit(ctx, msg, conf.CommitArgs); err != nil {
				return err
			}

			slog.Info("changes committed", "msg", msg)

			return nil
		},
		Short: "Commit your work",
		Use:   "commit",
	}

	return cmd
}
