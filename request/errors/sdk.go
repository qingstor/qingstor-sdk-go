package errors

import (
	"fmt"
)

// SDKError stores information of an error return by sdk itself.
type SDKError struct {
	Action    string
	RequestID string
	Err       error
}

// Unwrap implement interface of errors
func (e SDKError) Unwrap() error {
	return e.Err
}

// Error implement errors.Error
func (e SDKError) Error() string {
	return fmt.Sprintf("SDK Error: Action \"%s\", RequestID \"%s\", Err \"%s\"", e.Action, e.RequestID, e.Err)
}

// NewSDKError conduct a SDKError
func NewSDKError(fs ...func(*SDKError)) SDKError {
	e := SDKError{}
	for _, f := range fs {
		f(&e)
	}
	return e
}

// WithAction set Action for SDKError
func WithAction(action string) func(*SDKError) {
	return func(e *SDKError) {
		e.Action = action
	}
}

// WithError set Err for SDKError
func WithError(err error) func(*SDKError) {
	return func(e *SDKError) {
		e.Err = err
	}
}

// WithRequestID set RequestID for SDKError
func WithRequestID(id string) func(*SDKError) {
	return func(e *SDKError) {
		e.RequestID = id
	}
}
