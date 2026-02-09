package prompt_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pair/internal/prompt"
)

func TestPrefixSelect(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		selectFunc    prompt.PrefixSelectorFunc
		expected      string
		expectedError error
	}{
		"returns selected prefix": {
			selectFunc: func(opts []string, b bool) (string, error) {
				return "feat", nil
			},
			expected:      "feat",
			expectedError: nil,
		},
		"returns error when no prefix selected": {
			selectFunc: func(opts []string, b bool) (string, error) {
				return "", nil
			},
			expected:      "",
			expectedError: prompt.ErrNoPrefixSelected,
		},
		"returns error when selector errors": {
			selectFunc: func(opts []string, b bool) (string, error) {
				return "", errors.New("forced error")
			},
			expected:      "",
			expectedError: prompt.ErrSelectingPrefix,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			prefixSelector := prompt.PrefixSelector{
				SelectFunc: tc.selectFunc,
				Opts:       []string{"feat", "fix", "docs", "refactor"},
			}

			actual, err := prefixSelector.Select()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
