// package tuple provides tuple data types eg immutable sequence of items
package tuple

type Pair[T any, U any] struct {
	fst T
	snd U
}

func (p *Pair[T, U]) Fst() T {
	return p.fst
}

func (p *Pair[T, U]) Snd() U {
	return p.snd
}

func NewPair[T any, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		fst: first,
		snd: second,
	}
}
