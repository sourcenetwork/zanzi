package errors

import (
	"errors"
	"fmt"
	"strings"
)

var _ error = (*Error)(nil)

// NewWithCause returns a new zanzi error which wraps an ErrorKind and an underlying cause error
func NewWithCause(msg string, kind ErrorKind, cause error, pairs ...AttrPair) error {
	return &Error{
		Kind:     kind,
		Message:  msg,
		Cause:    cause,
		Metadata: appendPairs(nil, pairs),
	}
}

// New returns a new Zanzi base error type with the given kind
func New(message string, kind ErrorKind, pairs ...AttrPair) error {
	return &Error{
		Kind:     kind,
		Message:  message,
		Cause:    nil,
		Metadata: appendPairs(nil, pairs),
	}
}

// Error models a Zanzi error object, which may wrap an underlaying cause error
// and contains a set of string key-value pairs which contain request specific metadata
// which caused the error
type Error struct {
	Kind     ErrorKind
	Message  string
	Cause    error
	Metadata map[string]string
}

func (e *Error) Error() string {
	str := fmt.Sprintf("%v; data=%v; kind=%v", e.Message, e.Metadata, e.Kind.Error())
	if e.Cause != nil {
		str += "; cause: " + e.Cause.Error()
	}
	return str
}

// Is return true if target is of type ErrorKind and e.Kind is the same as ErrorKind
// or if target is also an Error instance and e's message contains target's message.
func (e *Error) Is(target error) bool {
	switch other := target.(type) {
	case *Error:
		return strings.Contains(e.Message, other.Message) && e.Kind == other.Kind
	case ErrorKind:
		return e.Kind.Is(other)
	default:
		return errors.Is(e.Cause, target) //not sure this is right
	}
}
