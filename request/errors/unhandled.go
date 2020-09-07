package errors

import (
	"fmt"
	"net/http"
)

// UnhandledResponseError stores information of an unhandled error response.
type UnhandledResponseError struct {
	StatusCode int
	Header     http.Header
	Content    string
}

// Error returns the description of QingStor error response.
func (e UnhandledResponseError) Error() string {
	return fmt.Sprintf(
		"Unhandled Error: StatusCode \"%d\", Header \"%v\", Content \"%s\"",
		e.StatusCode, e.Header, e.Content)
}

// NewUnhandledResponseError conduct an unhandled response error
func NewUnhandledResponseError(fs ...func(*UnhandledResponseError)) UnhandledResponseError {
	e := UnhandledResponseError{}
	for _, f := range fs {
		f(&e)
	}
	return e
}

// WithStatusCode set StatusCode for UnhandledResponseError
func WithStatusCode(code int) func(*UnhandledResponseError) {
	return func(e *UnhandledResponseError) {
		e.StatusCode = code
	}
}

// WithContent set Content for UnhandledResponseError
func WithContent(detail string) func(*UnhandledResponseError) {
	return func(e *UnhandledResponseError) {
		e.Content = detail
	}
}

// WithHeader set Header for UnhandledResponseError
func WithHeader(h http.Header) func(*UnhandledResponseError) {
	return func(e *UnhandledResponseError) {
		e.Header = h
	}
}
