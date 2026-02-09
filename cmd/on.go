package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
)

func NewCmdOn() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			var ticketID string

			if len(args) == 0 {
				ticketID = "prompt"
			} else {
				ticketID = args[0]
			}

			conf, err := config.Get()
			if err != nil {
				return err
			}

			if conf.TicketPrefix != "" {
				ticketID = conf.TicketPrefix + ticketID
			}

			// TODO: write to tmp file.
			// file path /tmp/pair/DATE/on
			log.Printf("ticket: %s", ticketID)

			return nil
		},
		Short: "Specify the ticket you are pairing on",
		Use:   "on",
	}

	return cmd
}
