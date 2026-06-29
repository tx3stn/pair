package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdAdd(_ *config.Config) *cobra.Command {
	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, file, err := config.Get()
			if err != nil {
				return err
			}

			coAuthor, err := prompt.AddCoAuthor(conf.SuggestedCoAuthorEmail, conf.AccessibleMode)
			if err != nil {
				return err
			}

			if conf.CoAuthors == nil {
				conf.CoAuthors = map[string]string{}
			}

			conf.CoAuthors[coAuthor.Name] = coAuthor.Email

			if err := config.Save(file, conf); err != nil {
				return err
			}

			slog.Info("added co-author", "name", coAuthor.Name, "email", coAuthor.Email)

			return nil
		},
		Short: "add a co-author to your config file",
		Use:   "add",
	}
}
