package cmd

import (
	"log/slog"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdWith(conf *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir, time.Now())

			_, err := setCoAuthors(session, conf)
			if err != nil {
				return err
			}

			return nil
		},
		Short: "Select who you are pairing with",
		Use:   "with",
	}

	return cmd
}

func setCoAuthors(session pairing.Session, conf *config.Config) ([]git.CoAuthor, error) {
	coAuthors := prompt.NewCoAuthorSelector(conf.CoAuthors, conf.AccessibleMode)

	selected, err := coAuthors.Select()
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
