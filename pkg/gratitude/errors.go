package gratitude

type GratitudeError struct {
	code    int
	message string
}

func (e *GratitudeError) Error() string {
	return e.message
}

func (e *GratitudeError) Status() int {
	return e.code
}

func NewGratitudeError(status int, msg string) error {
	return &GratitudeError{
		code:    status,
		message: msg,
	}
}
