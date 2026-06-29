package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pair/internal/config"
)

func TestFindConfigFile(t *testing.T) {
	testCases := map[string]struct {
		xdgEnvValue   string
		homeEnvValue  string
		expected      string
		expectedError error
	}{
		"ReturnsXdgFileWhenExists": {
			xdgEnvValue:   "testdata/xdg/valid",
			homeEnvValue:  "testdata/home/",
			expected:      "testdata/xdg/valid/pair.json",
			expectedError: nil,
		},
		"ReturnsHomeFileWhenExists": {
			xdgEnvValue:   "",
			homeEnvValue:  "testdata/home/",
			expected:      "testdata/home/.config/pair.json",
			expectedError: nil,
		},
		"ReturnsEmptyStringWhenNoEnvVarsAreSet": {
			xdgEnvValue:   "",
			homeEnvValue:  "",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)
			t.Setenv("HOME", tc.homeEnvValue)

			file, err := config.FindConfigFile()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, file)
		})
	}
}

func TestGet(t *testing.T) {
	testCases := map[string]struct {
		xdgEnvValue   string
		expectedError error
		expected      *config.Config
	}{
		"ReturnsErrorWhenFileIsInvalid": {
			xdgEnvValue:   "testdata/xdg/invalid",
			expectedError: config.ErrUnmashallingJSON,
			expected:      &config.Config{},
		},
		"ReturnsFileValidFileAsConfig": {
			xdgEnvValue:   "testdata/xdg/valid",
			expectedError: nil,
			expected: &config.Config{
				CoAuthors: map[string]string{
					"billy bob": "billy@billybob.org",
				},
				Prefixes: []string{
					"feat",
					"fix",
					"docs",
				},
				TicketPrefix: "JIRA-",
			},
		},
		"ReturnsErrorIfFileIsNotFound": {
			xdgEnvValue:   "testdata/xdg/missing",
			expectedError: config.ErrConfigNotFound,
			expected:      &config.Config{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Setenv("HOME", "")
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)

			actual, _, err := config.Get()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSave(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		conf     *config.Config
		expected string
	}{
		"writes all fields including schema and suggested email": {
			conf: &config.Config{
				Schema: "./schema.json",
				CoAuthors: map[string]string{
					"Jane Doe": "jane.doe@example.com",
				},
				Prefixes:               []string{"fix", "feat"},
				SuggestedCoAuthorEmail: "{{.FirstName}}.{{.LastName}}@example.com",
				TicketPrefix:           "TICKET-",
			},
			expected: `{
	"$schema": "./schema.json",
	"accessible": false,
	"coAuthors": {
		"Jane Doe": "jane.doe@example.com"
	},
	"commitArgs": "",
	"prefixes": [
		"fix",
		"feat"
	],
	"suggestedCoAuthorEmail": "{{.FirstName}}.{{.LastName}}@example.com",
	"ticketPrefix": "TICKET-"
}
`,
		},
		"omits empty optional fields": {
			conf: &config.Config{
				CoAuthors: map[string]string{
					"Jane Doe": "jane.doe@example.com",
				},
				Prefixes:     []string{"fix"},
				TicketPrefix: "TICKET-",
			},
			expected: `{
	"accessible": false,
	"coAuthors": {
		"Jane Doe": "jane.doe@example.com"
	},
	"commitArgs": "",
	"prefixes": [
		"fix"
	],
	"ticketPrefix": "TICKET-"
}
`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			file := filepath.Join(t.TempDir(), "pair.json")

			require.NoError(t, config.Save(file, tc.conf))

			content, err := os.ReadFile(filepath.Clean(file))
			require.NoError(t, err)
			assert.Equal(t, tc.expected, string(content))
		})
	}
}
