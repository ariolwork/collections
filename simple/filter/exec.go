package filter

func SliceBy[T any](source []T, filter func(T) bool) []T {
	result := make([]T, 0, len(source))

	for _, item := range source {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func MapBy[K comparable, V any](source map[K]V, filter func(K, V) bool) map[K]V {
	result := make(map[K]V, len(source))

	for key, value := range source {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}
