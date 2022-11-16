package types

// Option type is a container for a value
// Option may or may not contain value
type Option[T any] struct {
    val T
    empty bool
}

// Return a copy of the options type
// Calling Value on an empty option will panic
func (o *Option[T]) Value() T {
    if o.IsEmpty() {
        panic("cannot return value for empty Option")
    }
    return o.val
}

// Return true if Option does not contain a value
func (o *Option[T]) IsEmpty() bool {
    return o.empty
}

// Build an Option with val as its inner value
func Some[T any](val T) Option[T] {
    return Option[T] {
        val: val,
        empty: false,
    }
}

// Build an empty Option
func None[T any]() Option[T] {
    return Option[T] {
        empty: true,
    }
}
