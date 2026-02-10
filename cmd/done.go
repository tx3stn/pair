package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pair/internal/pairing"
)

func NewCmdDone() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			session := pairing.NewSession(pairing.DataDir, time.Now())

			if err := session.Clean(); err != nil {
				return err
			}

			log.Println("pairing session complete!")

			return nil
		},
		Short: "All done - remove the active pairing state",
		Use:   "done",
	}

	return cmd
}
