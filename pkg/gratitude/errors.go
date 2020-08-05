package gratitude

// Error is a custom error type for gratitude service operations
type Error struct {
	code    int
	message string
}

// Error returns the string contained in error.message
func (e *Error) Error() string {
	return e.message
}

// Status returns the status code contained in error.code
func (e *Error) Status() int {
	return e.code
}

// NewGratitudeError is a constructor for the error type in the gratitude package
func NewGratitudeError(status int, msg string) error {
	return &Error{
		code:    status,
		message: msg,
	}
}
