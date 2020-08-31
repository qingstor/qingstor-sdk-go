package errors

import "fmt"

// UnhandledError stores information of an unhandled error response.
type UnhandledError struct {
	StatusCode int

	Message   string
	Detail    string
	RequestID string
}

// Error returns the description of QingStor error response.
func (e UnhandledError) Error() string {
	return fmt.Sprintf(
		"Unhandled Error: StatusCode \"%d\", Message \"%s\", Request ID \"%s\", Detail \"%s\"",
		e.StatusCode, e.Message, e.RequestID, e.Detail)
}

// NewUnhandledError conduct an unhandled error
func NewUnhandledError(fs ...func(*UnhandledError)) UnhandledError {
	e := UnhandledError{}
	for _, f := range fs {
		f(&e)
	}
	return e
}

// WithRequestID set RequestID for UnhandledError
func WithRequestID(id string) func(*UnhandledError) {
	return func(e *UnhandledError) {
		e.RequestID = id
	}
}

// WithStatusCode set StatusCode for UnhandledError
func WithStatusCode(code int) func(*UnhandledError) {
	return func(e *UnhandledError) {
		e.StatusCode = code
	}
}

// WithDetail set Detail for UnhandledError
func WithDetail(detail string) func(*UnhandledError) {
	return func(e *UnhandledError) {
		e.Detail = detail
	}
}

// WithMessage set Message for UnhandledError
func WithMessage(msg string) func(*UnhandledError) {
	return func(e *UnhandledError) {
		e.Message = msg
	}
}
