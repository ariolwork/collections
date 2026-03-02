package slices

func ToMap[S ~[]T, T any, K comparable](seq S, keySelector func(T) K) map[K]T {
	result := make(map[K]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = val
	}

	return result
}

func ToMapBuckets[S ~[]T, T any, K comparable](seq S, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = append(result[keySelector(val)], val)
	}

	return result
}

func Transform[S ~[]T, T, R any](source S, applyFunc func(T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, applyFunc(item))
	}

	return result
}
