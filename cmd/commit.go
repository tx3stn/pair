package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/git"
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

			// Read ticket id
			// Read coauthors
			// Create text area with coauthors and prefix to type message
			commitMsg := selected + "(): "

			if err := git.Commit(commitMsg, conf.CommitArgs); err != nil {
				return err
			}

			log.Printf("committed: %s", commitMsg)

			return nil
		},
		Short: "Commit your work",
		Use:   "commit",
	}

	return cmd
}
