package types

var _ Logger = (*NoopLogger)(nil)

// Logger defines a generic facade for a log service.
// A concrete implementation is given during lib setup by the caller.
// This gives callers flexibility by not tieing zanzi to a specific
// logging framework.
type Logger interface {
	Errorf(msg string, args ...any)
	Warnf(msg string, args ...any)
	Infof(msg string, args ...any)
	Debugf(msg string, args ...any)
}

// NoopLogger does nothing when called
type NoopLogger struct{}

func (l *NoopLogger) Errorf(msg string, args ...any) {}
func (l *NoopLogger) Warnf(msg string, args ...any)  {}
func (l *NoopLogger) Infof(msg string, args ...any)  {}
func (l *NoopLogger) Debugf(msg string, args ...any) {}
