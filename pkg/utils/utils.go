package utils

func MapSlice[T any, U any](ts []T, mapper func(T) U) []U {
	us := make([]U, 0, len(ts))

	for _, t := range ts {
		u := mapper(t)
		us = append(us, u)
	}

	return us
}

func ConsumeChan[T any](ch <-chan T) []T {
	var values []T
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		values = append(values, val)
	}
	return values
}
