package pairing_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pair/internal/git"
	"github.com/tx3stn/pair/internal/pairing"
)

func TestSessionGetCoAuthors(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		sessionDir    string
		expected      []git.CoAuthor
		expectedError error
	}{
		"returns co-authors when file exists": {
			sessionDir: "./testdata",
			expected: []git.CoAuthor{
				{Name: "alice", Email: "alice@example.com"},
				{Name: "bob", Email: "bob@example.com"},
			},
			expectedError: nil,
		},
		"returns empty slice when file does not exist": {
			sessionDir:    "./foo",
			expected:      []git.CoAuthor{},
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			session := pairing.NewSession(tc.sessionDir)

			actual, err := session.GetCoAuthors()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSessionSetCoAuthors(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		coAuthors     []git.CoAuthor
		expected      string
		expectedError error
	}{
		"sets co-authors successfully": {
			coAuthors: []git.CoAuthor{
				{Name: "alice", Email: "alice@example.com"},
				{Name: "bob", Email: "bob@example.com"},
			},
			expected: `{"name":"alice","email":"alice@example.com"}
{"name":"bob","email":"bob@example.com"}`,
			expectedError: nil,
		},
		"sets empty co-authors successfully": {
			coAuthors:     []git.CoAuthor{},
			expected:      "",
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			session := pairing.NewSession(t.TempDir())
			err := session.SetCoAuthors(tc.coAuthors)
			require.ErrorIs(t, err, tc.expectedError)

			if tc.expectedError != nil {
				return
			}

			content, readErr := os.ReadFile(session.WithFile)
			require.NoError(t, readErr)
			assert.Equal(t, tc.expected, string(content))
		})
	}
}

func TestSessionGetTicketID(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		sessionDir    string
		expected      string
		expectedError error
	}{
		"returns ticket ID when file exists": {
			sessionDir:    "./testdata",
			expected:      "TICKET-123",
			expectedError: nil,
		},
		"returns empty string when file does not exist": {
			sessionDir:    "./foo",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			session := pairing.NewSession(tc.sessionDir)

			actual, err := session.GetTicketID()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSessionSetTicketID(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ticketID      string
		expectedError error
	}{
		"sets ticket ID successfully": {
			ticketID:      "TICKET-123",
			expectedError: nil,
		},
		"sets empty ticket ID successfully": {
			ticketID:      "",
			expectedError: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			session := pairing.NewSession(t.TempDir())
			err := session.SetTicketID(tc.ticketID)
			require.ErrorIs(t, err, tc.expectedError)

			if tc.expectedError != nil {
				return
			}

			content, readErr := os.ReadFile(session.OnFile)
			require.NoError(t, readErr)
			assert.Equal(t, tc.ticketID, string(content))
		})
	}
}
