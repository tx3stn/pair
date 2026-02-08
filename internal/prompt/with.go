// Package prompt contains logic related to user interactions.
package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pair/internal/git"
)

type CoAuthorSelectorFunc func(map[string]string, bool) ([]string, error)

type CoAuthorSelector struct {
	SelectFunc     CoAuthorSelectorFunc
	Opts           map[string]string
	accessibleMode bool
}

func NewCoAuthorSelector(opts map[string]string, accessible bool) CoAuthorSelector {
	return CoAuthorSelector{
		SelectFunc:     selectCoAuthors,
		Opts:           opts,
		accessibleMode: accessible,
	}
}

func (c CoAuthorSelector) Select() ([]git.CoAuthor, error) {
	selected, err := c.SelectFunc(c.Opts, c.accessibleMode)
	if err != nil {
		return []git.CoAuthor{}, fmt.Errorf("%w: %w", ErrSelectingCoAuthors, err)
	}

	if len(selected) == 0 {
		return []git.CoAuthor{}, ErrNoCoAuthorsSelected
	}

	coAuthors := make([]git.CoAuthor, len(selected))
	for i, name := range selected {
		coAuthors[i] = git.CoAuthor{Name: name, Email: c.Opts[name]}
	}

	return coAuthors, nil
}

//nolint:wrapcheck
func selectCoAuthors(opts map[string]string, accessible bool) ([]string, error) {
	var selected []string

	options := make([]huh.Option[string], 0, len(opts))

	for name := range opts {
		options = append(options, huh.NewOption(name, name))
	}

	prompt := huh.NewMultiSelect[string]().
		Title("paring with:").
		Options(options...).
		Value(&selected)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return []string{}, err
		}
	} else {
		if err := prompt.Run(); err != nil {
			return []string{}, err
		}
	}

	return selected, nil
}
