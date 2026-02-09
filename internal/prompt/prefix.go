package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

type PrefixSelectorFunc func([]string, bool) (string, error)

type PrefixSelector struct {
	SelectFunc     PrefixSelectorFunc
	Opts           []string
	accessibleMode bool
}

func NewPrefixSelector(opts []string, accessible bool) PrefixSelector {
	return PrefixSelector{
		SelectFunc:     selectPrefix,
		Opts:           opts,
		accessibleMode: accessible,
	}
}

func (p PrefixSelector) Select() (string, error) {
	selected, err := p.SelectFunc(p.Opts, p.accessibleMode)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingPrefix, err)
	}

	if selected == "" {
		return "", ErrNoPrefixSelected
	}

	return selected, nil
}

//nolint:wrapcheck
func selectPrefix(opts []string, accessible bool) (string, error) {
	var selected string

	options := make([]huh.Option[string], 0, len(opts))

	for _, prefix := range opts {
		options = append(options, huh.NewOption(prefix, prefix))
	}

	prompt := huh.NewSelect[string]().
		Title("select prefix:").
		Options(options...).
		Value(&selected)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", err
		}
	} else {
		if err := prompt.Run(); err != nil {
			return "", err
		}
	}

	return selected, nil
}
