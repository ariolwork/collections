package simple

type MapRecord[K comparable, V any] struct {
	Key   K
	Value V
}

func SliceToMap[T any, K comparable](seq []T, keySelector func(T) K) map[K]T {
	result := make(map[K]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = val
	}

	return result
}

func SliceToMapBuckets[T any, K comparable](seq []T, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = append(result[keySelector(val)], val)
	}

	return result
}

func MapToSlice[T any, K comparable](seq map[K]T, keySelector func(T) K) []MapRecord[K, T] {
	result := make([]MapRecord[K, T], 0, len(seq))

	for key, val := range seq {
		result = append(result, MapRecord[K, T]{key, val})
	}

	return result
}
