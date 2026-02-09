package prompt

// Error is the error type.
type Error uint8

const (
	// ErrSelectingCoAuthors is the error returned when selecting co authors fails.
	ErrSelectingCoAuthors Error = iota + 1
	// ErrNoCoAuthorsSelected is the error returned when you don't select any co authors
	// in the prompt.
	ErrNoCoAuthorsSelected
	// ErrSelectingPrefix is the error returned when selecting prefix fails.
	ErrSelectingPrefix
	// ErrNoPrefixSelected is the error returned when you don't select any prefix
	// in the prompt.
	ErrNoPrefixSelected
)

// Error returns the message string for the given error.
func (e Error) Error() string {
	switch e {
	case ErrSelectingCoAuthors:
		return "error selecting co-authors"
	case ErrNoCoAuthorsSelected:
		return "no co-authors selected"
	case ErrSelectingPrefix:
		return "error selecting prefix"
	case ErrNoPrefixSelected:
		return "no prefix selected"
	default:
		return "unknown error"
	}
}
