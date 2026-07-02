// Package config contains logic related to user config files.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the options defined in the config file.
type Config struct {
	// Schema is the optional "$schema" key used for in-editor validation. It is
	// modelled here so it is preserved when the config is written back to disk.
	Schema                 string            `json:"$schema,omitempty"`
	AccessibleMode         bool              `json:"accessible"`
	CoAuthors              map[string]string `json:"coAuthors"`
	CommitArgs             string            `json:"commitArgs"`
	Prefixes               []string          `json:"prefixes"`
	SuggestedCoAuthorEmail string            `json:"suggestedCoAuthorEmail,omitempty"`
	TicketPrefix           string            `json:"ticketPrefix"`
}

// Get returns the config read from the file.
func Get() (*Config, string, error) {
	file, err := FindConfigFile()
	if err != nil {
		return &Config{}, "", fmt.Errorf("error checking for existence of config file: %w", err)
	}

	if file == "" {
		return &Config{}, "", ErrConfigNotFound
	}

	content, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return &Config{}, "", fmt.Errorf("%w: %w", ErrReadingConfigFile, err)
	}

	var conf Config
	if err = json.Unmarshal(content, &conf); err != nil {
		return &Config{}, "", fmt.Errorf("%w: %w", ErrUnmashallingJSON, err)
	}

	return &conf, file, nil
}

// FindConfigFile checks the expected paths for a pair.json config file and returns
// the path to it if found.
// The paths are checked in the order of precedence:
//   - XDG_CONFIG_DIR
//   - HOME/.config
func FindConfigFile() (string, error) {
	paths := []string{}
	configFileName := "pair.json"

	if xdg, ok := os.LookupEnv("XDG_CONFIG_DIR"); ok {
		paths = append(paths, xdg)
	}

	if home, ok := os.LookupEnv("HOME"); ok {
		paths = append(paths, filepath.Join(home, ".config"))
	}

	if len(paths) == 0 {
		return "", nil
	}

	for _, path := range paths {
		file := filepath.Join(path, configFileName)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			// no config file at location, continue looking.
			continue
		}

		return file, nil
	}

	return "", nil
}

// Save writes the config back to the given file as tab-indented JSON.
func Save(file string, conf *Config) error {
	content, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMarshallingJSON, err)
	}

	content = append(content, '\n')

	if err := os.WriteFile(filepath.Clean(file), content, 0o600); err != nil {
		return fmt.Errorf("%w: %w", ErrWritingConfigFile, err)
	}

	return nil
}
