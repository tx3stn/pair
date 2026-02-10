package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/config"
	"github.com/tx3stn/pair/internal/pairing"
	"github.com/tx3stn/pair/internal/prompt"
)

func NewCmdOn() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return err
			}

			session := pairing.NewSession(pairing.DataDir, time.Now())

			ticketID := ""

			if len(args) > 0 {
				ticketID = args[0]
			}

			_, err = setTicketID(session, conf, ticketID)

			return err
		},
		Short: "Specify the ticket you are pairing on",
		Use:   "on",
	}

	return cmd
}

func setTicketID(session pairing.Session, conf config.Config, ticketID string) (string, error) {
	var err error

	if ticketID == "" {
		ticketID, err = prompt.TicketID(conf.TicketPrefix, conf.AccessibleMode)
		if err != nil {
			return "", err
		}
	}

	if err := session.SetTicketID(ticketID); err != nil {
		return "", err
	}

	// TODO: debug log output
	log.Printf("set ticket id: %s", ticketID)

	return ticketID, nil
}
