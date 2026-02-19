package sliceops

func Filter[T any](source []T, filter func(T) bool) []T {
	result := make([]T, 0, len(source))

	for _, item := range source {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}
