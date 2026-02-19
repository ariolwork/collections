package slices_exp

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

func Transform[T, R any](source []T, applyFunc func(T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, applyFunc(item))
	}

	return result
}
