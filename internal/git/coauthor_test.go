package git_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pair/internal/git"
)

func TestSuggestEmail(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		format        string
		firstName     string
		lastName      string
		expected      string
		expectedError error
	}{
		"empty format returns empty string": {
			format:    "",
			firstName: "Jane",
			lastName:  "Doe",
			expected:  "",
		},
		"first name and last name": {
			format:    "{{.FirstName}}.{{.LastName}}@company.com",
			firstName: "Jane",
			lastName:  "Doe",
			expected:  "jane.doe@company.com",
		},
		"first initial and last name": {
			format:    "{{.FirstInitial}}{{.LastName}}@company.com",
			firstName: "Jane",
			lastName:  "Doe",
			expected:  "jdoe@company.com",
		},
		"last initial": {
			format:    "{{.FirstName}}{{.LastInitial}}@company.com",
			firstName: "Jane",
			lastName:  "Doe",
			expected:  "janed@company.com",
		},
		"result is lowercased": {
			format:    "{{.FirstName}}.{{.LastName}}@Company.com",
			firstName: "JANE",
			lastName:  "McDoe",
			expected:  "jane.mcdoe@company.com",
		},
		"empty last name does not panic": {
			format:    "{{.FirstInitial}}{{.LastName}}@company.com",
			firstName: "Jane",
			lastName:  "",
			expected:  "j@company.com",
		},
		"multi-byte initial uses the rune": {
			format:    "{{.FirstInitial}}{{.LastName}}@company.com",
			firstName: "Élise",
			lastName:  "Adams",
			expected:  "éadams@company.com",
		},
		"invalid template returns error": {
			format:        "{{.FirstName",
			firstName:     "Jane",
			lastName:      "Doe",
			expectedError: git.ErrSuggestingEmail,
		},
		"unknown field returns error": {
			format:        "{{.Middle}}@company.com",
			firstName:     "Jane",
			lastName:      "Doe",
			expectedError: git.ErrSuggestingEmail,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := git.SuggestEmail(tc.format, tc.firstName, tc.lastName)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
