package utils

func MapSlice[T any, U any](ts []T, mapper func(T) U) []U {
	us := make([]U, 0, len(ts))

	for _, t := range ts {
		u := mapper(t)
		us = append(us, u)
	}

	return us
}

func MapSliceErr[T any, U any](ts []T, mapper func(T) (U, error)) ([]U, error) {
	us := make([]U, 0, len(ts))

	for _, t := range ts {
		u, err := mapper(t)
                if err != nil {
                    return nil, err
                }
		us = append(us, u)
	}

	return us, nil
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

func Conditional[T any](proposition bool, trueChoice, falseChoice T) T {
    if proposition {
        return trueChoice
    } else {
        return falseChoice
    }
}

func Filter[T any](ts []T, filter func (T) bool) []T {
    filtered := make([]T, 0, len(ts))
    for _, t := range ts {
        if filter(t) {
            filtered = append(filtered, t)
        }
    }
    return filtered
}

func Identity[T any](t T) T {
    return t
}
