package prompt

import (
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func EditCommitMessage(message string, coAuthors []string, accessible bool) (string, error) {
	edited := message

	description := ""
	if len(coAuthors) > 0 {
		description = strings.Join(coAuthors, "\n")
	}

	prompt := huh.NewText().
		Title("edit commit message:").
		Description(description).
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

	// Add co-authors back to the final message
	if len(coAuthors) > 0 {
		edited += "\n\n" + strings.Join(coAuthors, "\n")
	}

	return edited, nil
}
