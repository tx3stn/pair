package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdCommit(conf *config.Config) *cobra.Command {
	short := "commit your work"

	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			prefix := prompt.NewPrefixSelector(conf.Prefixes, conf.AccessibleMode)

			commitType, err := prefix.Select()
			if err != nil {
				return err
			}

			session := pairing.NewSession(pairing.DataDir)

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

			if len(coAuthors) == 0 || (len(args) > 0 && args[0] == "+") {
				coAuthors, err = setCoAuthors(session, conf, true)
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
		Short: short,
		Use:   "commit",
		Long: short + `

if you already have an active pairing session those values will be
automatically used

if you don't, you will be prompted to enter ticket id and select co-authors

to add an additional co-author to the current selection pass the '+' arg, e.g.:

pair commit +
`,
	}
}
