package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pair/internal/git"
)

// AddCoAuthor prompts for a co-author's first name, last name and email and
// returns the resulting co-author.
//
// The email input is pre-filled with a suggestion rendered from emailFormat
// using the entered names, which the user can accept or edit.
// An empty emailFormat means the email input starts blank.
func AddCoAuthor(emailFormat string, accessible bool) (git.CoAuthor, error) {
	var firstName, lastName string

	firstNameInput := huh.NewInput().
		Title("first name:").
		Validate(required).
		Value(&firstName)

	if err := runInput(firstNameInput, accessible); err != nil {
		return git.CoAuthor{}, ErrPromptingCoAuthor
	}

	lastNameInput := huh.NewInput().
		Title("last name:").
		Value(&lastName)

	if err := runInput(lastNameInput, accessible); err != nil {
		return git.CoAuthor{}, ErrPromptingCoAuthor
	}

	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	email, err := git.SuggestEmail(emailFormat, firstName, lastName)
	if err != nil {
		return git.CoAuthor{}, fmt.Errorf("error suggesting coauthor email: %w", err)
	}

	emailInput := huh.NewInput().
		Title("email:").
		Value(&email)

	// Only require a value when there is no suggestion to fall back to. With a
	// suggestion present, submitting an empty value accepts the suggestion (in
	// accessible mode the input validates the typed value, not the default).
	if email == "" {
		emailInput = emailInput.Validate(required)
	}

	if err := runInput(emailInput, accessible); err != nil {
		return git.CoAuthor{}, ErrPromptingCoAuthor
	}

	name := strings.TrimSpace(firstName + " " + lastName)

	return git.CoAuthor{Name: name, Email: strings.TrimSpace(email)}, nil
}

func runInput(input *huh.Input, accessible bool) error {
	if accessible {
		//nolint:wrapcheck
		return input.RunAccessible(os.Stdout, os.Stdin)
	}

	//nolint:wrapcheck
	return input.Run()
}

func required(value string) error {
	if strings.TrimSpace(value) == "" {
		return ErrEmptyValue
	}

	return nil
}
