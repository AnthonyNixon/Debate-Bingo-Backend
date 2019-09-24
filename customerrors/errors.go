package customerrors

type Error interface {
	StatusCode() int
	Description() string
}

// New returns an error that formats as the given text.
func New(status int, description string) Error {
	return &customError{status, description}
}

// errorString is a trivial implementation of error.
type customError struct {
	status int
	description string
}

func (e *customError) Description() string {
	return e.description
}

func (e *customError) StatusCode() int {
	return e.status
}