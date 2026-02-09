package prompt

import (
	"os"

	"github.com/charmbracelet/huh"
)

func EditCommitMessage(message string, accessible bool) (string, error) {
	edited := message

	prompt := huh.NewText().
		Title("commit message:").
		Value(&edited)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", ErrEditingCommitMessage
		}
	} else {
		if err := prompt.Run(); err != nil {
			return "", ErrEditingCommitMessage
		}
	}

	return edited, nil
}
