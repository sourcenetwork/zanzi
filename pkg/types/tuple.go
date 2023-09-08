package types

// Pair models as 2-tuple
type Pair[T any, U any] struct {
	first T
	second U
}

func (p *Pair[T, U]) First() T {
	return p.first
}

func (p *Pair[T, U]) Second() U {
	return p.second
}

func NewPair[T any, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		first: first,
		second: second,
	}
}
