package errors

import "errors"

var _ error = ErrorKind(1)

// ErrorKind models a high level grouping of errors
type ErrorKind int

const (
	// Unknown models an error whose Kind wasn't initialized
	Unknown ErrorKind = 1
	// Internal: an internal error condition happened
	Internal ErrorKind = 2
	// BadInput: cannot process the request with the given input
	BadInput ErrorKind = 3
	// NotFound: the requested resource was not found
	NotFound ErrorKind = 4
	// Unauthorized: the authenticated user does not have permission to execute the operation
	Unauthorized ErrorKind = 5
	// Violation: a codition which would compromise a system invariant if executed
	Violation ErrorKind = 6
)

func (k ErrorKind) String() string {
	switch k {
	case Internal:
		return "internal error"
	case BadInput:
		return "bad input"
	case NotFound:
		return "not found"
	case Unauthorized:
		return "unauthorized"
	case Violation:
		return "logic violation"
	default:
		return "unknown"
	}
}

func (k ErrorKind) Error() string {
	return k.String()
}

func (k ErrorKind) Is(target error) bool {
	switch t := target.(type) {
	case ErrorKind:
		return int(k) == int(t)
	default:
		return errors.Is(k, target)
	}
}
