package errors

// AttrPair model a key value attribute pair
type AttrPair struct {
	Key   string
	Value string
}

// Pair returns a new AttrPair with the given key value
func Pair(key, value string) AttrPair {
	return AttrPair{
		Key:   key,
		Value: value,
	}
}

// Wrap attenuates an error type by adding an extra error message
// as well as a set of key-value metadata pairs
func Wrap(msg string, err error, pairs ...AttrPair) error {
	switch e := err.(type) {
	case ErrorKind:
		return &Error{
			Kind:     e,
			Message:  msg,
			Cause:    nil,
			Metadata: appendPairs(nil, pairs),
		}
	case *Error:
		return &Error{
			Kind:     e.Kind,
			Message:  msg + ": " + e.Message,
			Cause:    e.Cause,
			Metadata: appendPairs(e.Metadata, pairs),
		}
	default:
		return &Error{
			Kind:     Unknown,
			Message:  msg,
			Cause:    err,
			Metadata: nil,
		}
	}
}

// Attrs attenuates an error by adding additional key-value context metadata.
// Overwrites keys with the same name
func Attrs(err error, pairs ...AttrPair) error {
	switch e := err.(type) {
	case ErrorKind:
		return &Error{
			Kind:     e,
			Message:  "",
			Cause:    nil,
			Metadata: appendPairs(nil, pairs),
		}
	case *Error:
		return &Error{
			Kind:     e.Kind,
			Message:  e.Message,
			Cause:    e.Cause,
			Metadata: appendPairs(e.Metadata, pairs),
		}
	default:
		return &Error{
			Kind:     Unknown,
			Message:  "",
			Cause:    err,
			Metadata: appendPairs(nil, pairs),
		}
	}
}

func appendPairs(m map[string]string, pairs []AttrPair) map[string]string {
	attrs := make(map[string]string, len(m)+len(pairs))
	for k, v := range m {
		attrs[k] = v
	}
	for _, p := range pairs {
		attrs[p.Key] = p.Value
	}
	return attrs
}
