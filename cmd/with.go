package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdWith() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return err
			}

			coAuthors := prompt.NewCoAuthorSelector(conf.CoAuthors, conf.AccessibleMode)

			selected, err := coAuthors.Select()
			if err != nil {
				return err
			}

			for _, co := range selected {
				log.Println(co.Format())
			}

			return nil
		},
		Short: "Select who you're pairing with",
		Use:   "with",
	}

	return cmd
}
