package prompt_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/prompt"
)

func TestCoAuthorSelect(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		selectFunc    prompt.CoAuthorSelectorFunc
		expected      []git.CoAuthor
		expectedError error
	}{
		"returns selected co-authors": {
			selectFunc: func(m map[string]string, b bool) ([]string, error) {
				return []string{"anne", "bob"}, nil
			},
			expected: []git.CoAuthor{
				{Name: "anne", Email: "anne@example.org"},
				{Name: "bob", Email: "bob@example.org"},
			},
			expectedError: nil,
		},
		"returns error when no co-authors selected": {
			selectFunc: func(m map[string]string, b bool) ([]string, error) {
				return []string{}, nil
			},
			expected:      []git.CoAuthor{},
			expectedError: prompt.ErrNoCoAuthorsSelected,
		},
		"returns error when selector errors": {
			selectFunc: func(m map[string]string, b bool) ([]string, error) {
				return []string{}, errors.New("forced error")
			},
			expected:      []git.CoAuthor{},
			expectedError: prompt.ErrSelectingCoAuthors,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			coAuthors := prompt.CoAuthorSelector{
				SelectFunc: tc.selectFunc,
				Opts: map[string]string{
					"anne":  "anne@example.org",
					"bob":   "bob@example.org",
					"carol": "carol@example.org",
				},
			}

			actual, err := coAuthors.Select()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
