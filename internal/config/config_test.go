package config_test

import (
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

			actual, err := config.Get()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
