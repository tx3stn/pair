package prompt

import (
	"os"

	"github.com/charmbracelet/huh"
)

func TicketID(prefix string, accessible bool) (string, error) {
	ticketID := prefix

	prompt := huh.NewInput().
		Title("enter ticket ID:").
		Value(&ticketID)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", ErrPromptingTicketID
		}
	} else {
		if err := prompt.Run(); err != nil {
			return "", ErrPromptingTicketID
		}
	}

	return ticketID, nil
}
