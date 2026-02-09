package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
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

			// TODO: debug log selected
			for _, co := range selected {
				log.Println(co.Format())
			}

			session := pairing.NewSession(pairing.DataDir, time.Now())
			if err := session.SetCoAuthors(selected); err != nil {
				return err
			}

			return nil
		},
		Short: "Select who you're pairing with",
		Use:   "with",
	}

	return cmd
}
