package errors

import (
	"errors"
	"fmt"
)

// SDKError stores information of an error return by sdk it self.
type SDKError struct {
	Action    string
	RequestID string
	Err       error
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

// WithErrStr set Err by string for SDKError
func WithErrStr(errStr string) func(*SDKError) {
	return func(e *SDKError) {
		e.Err = errors.New(errStr)
	}
}

// WithRequestID set RequestID for SDKError
func WithRequestID(id string) func(*SDKError) {
	return func(e *SDKError) {
		e.RequestID = id
	}
}
