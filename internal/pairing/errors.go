package pairing

// Error is the error type.
type Error uint8

const (
	// ErrReadingCoAuthors is the error returned when reading co-authors fails.
	ErrReadingCoAuthors Error = iota + 1
	// ErrWritingCoAuthors is the error returned when writing co-authors fails.
	ErrWritingCoAuthors
	// ErrReadingTicketID is the error returned when reading ticket ID fails.
	ErrReadingTicketID
	// ErrWritingTicketID is the error returned when writing ticket ID fails.
	ErrWritingTicketID
	// ErrCreatingDirectory is the error returned when creating directory fails.
	ErrCreatingDirectory
	// ErrMarshalingCoAuthor is the error returned when the co-author marshal fails.
	ErrMarshalingCoAuthor
	// ErrUnmarshalingCoAuthor is the error returned when the co-author unmarshal fails.
	ErrUnmarshalingCoAuthor
)

// Error returns the message string for the given error.
func (e Error) Error() string {
	switch e {
	case ErrReadingCoAuthors:
		return "error reading co-authors"
	case ErrWritingCoAuthors:
		return "error writing co-authors"
	case ErrReadingTicketID:
		return "error reading ticket ID"
	case ErrWritingTicketID:
		return "error writing ticket ID"
	case ErrCreatingDirectory:
		return "error creating directory"
	case ErrMarshalingCoAuthor:
		return "error marshaling co-author"
	case ErrUnmarshalingCoAuthor:
		return "error unmarshaling co-author"
	default:
		return "unknown error"
	}
}
