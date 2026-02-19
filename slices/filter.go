package slices

func IdempotentFilter[S ~[]T, T any](s S, filter func(T) bool) []T {
	result := make([]T, 0, len(s))

	for _, item := range s {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}
