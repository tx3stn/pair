package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pair/internal/git"
)

func EditCommitMessage(message string, coAuthors []git.CoAuthor, accessible bool) (string, error) {
	edited := message

	description := ""

	if len(coAuthors) > 0 {
		var output strings.Builder
		for _, coAuthor := range coAuthors {
			if _, err := output.WriteString(coAuthor.Format() + "\n"); err != nil {
				return "", fmt.Errorf("error formatting co-authors: %w", err)
			}
		}

		description = strings.TrimSuffix(output.String(), "\n")
	}

	prompt := huh.NewText().
		Title("edit commit message:").
		Description(description).
		Value(&edited)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", ErrEditingCommitMessage
		}

		// Accessible mode removes the prefix.
		edited = message + edited
	} else {
		if err := prompt.Run(); err != nil {
			return "", ErrEditingCommitMessage
		}
	}

	// Add co-authors back to the final message
	if len(coAuthors) > 0 {
		edited += "\n\n" + description
	}

	return edited, nil
}
