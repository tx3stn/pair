package git

// Error is the error type.
type Error uint8

const (
	// ErrRunningGitCommand is the error returned when running git commands fails.
	ErrRunningGitCommand Error = iota + 1
)

// Error returns the message string for the given error.
func (e Error) Error() string {
	switch e {
	case ErrRunningGitCommand:
		return "error running git commands"
	default:
		return "unknown error"
	}
}
