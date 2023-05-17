package shared

func Map[I, O any, S ~[]I](slice S, fn func(I) O) []O {
	result := make([]O, 0, len(slice))

	for _, entry := range slice {
		result = append(result, fn(entry))
	}

	return result
}

func MapCheck[I, O any, S ~[]I](slice S, fn func(I) (O, bool)) []O {
	result := make([]O, 0, len(slice))

	for _, entry := range slice {
		value, ok := fn(entry)
		if !ok {
			continue
		}

		result = append(result, value)
	}

	return result
}
