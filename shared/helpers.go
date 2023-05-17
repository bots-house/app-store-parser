package shared

import "strings"

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

func Filter[T any, S ~[]T](slice S, fn func(T) bool) S {
	result := make(S, 0, len(slice))

	for _, entry := range slice {
		if !fn(entry) {
			continue
		}

		result = append(result, entry)
	}

	return result
}

func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	result := make([]K, 0, len(m))

	for key := range m {
		result = append(result, key)
	}

	return result
}

func In[K comparable](entry K, items ...K) bool {
	for _, item := range items {
		if item == entry {
			return true
		}
	}

	return false
}

func GetCountryHeader(country, separator string) string {
	code, ok := countryMap[strings.ToLower(country)]
	if !ok {
		code = countryMap["us"]
	}

	if separator != "" {
		code += "," + separator
	}

	return code
}
