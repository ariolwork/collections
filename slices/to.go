package slices

func ToMap[T any, K comparable](seq []T, keySelector func(T) K) map[K]T {
	result := make(map[K]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = val
	}

	return result
}

func ToMapBuckets[T any, K comparable](seq []T, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = append(result[keySelector(val)], val)
	}

	return result
}
