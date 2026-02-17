package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdWith(conf *config.Config) *cobra.Command {
	short := "select who you are pairing with"

	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir)

			useSelected := len(args) == 0 || args[0] != "new"

			if _, err := setCoAuthors(session, conf, useSelected); err != nil {
				return err
			}

			return nil
		},
		Short: short,
		Use:   "with",
		Long: short + `

launches a multi select prompt to select co-authors from your config file

uses the currently selected co-authors from your session, to replace the 
co-authors completely pass the 'new' arg, e.g.:

pair with new
`,
	}
}

func setCoAuthors(
	session *pairing.Session,
	conf *config.Config,
	useSelected bool,
) ([]git.CoAuthor, error) {
	currentlySelected := []git.CoAuthor{}

	if useSelected {
		var err error

		currentlySelected, err = session.GetCoAuthors()
		if err != nil {
			return []git.CoAuthor{}, err
		}
	}

	coAuthors := prompt.NewCoAuthorSelector(conf.CoAuthors, conf.AccessibleMode)

	selected, err := coAuthors.Select(currentlySelected)
	if err != nil {
		return []git.CoAuthor{}, err
	}

	for _, co := range selected {
		slog.Debug("added co-author", "coAuthor", co.Format())
	}

	if err := session.SetCoAuthors(selected); err != nil {
		return []git.CoAuthor{}, err
	}

	return session.GetCoAuthors()
}
