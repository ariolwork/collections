package mapops

type MapRecord[K comparable, V any] struct {
	Key   K
	Value V
}

func ToSlice[T any, K comparable](seq map[K]T, keySelector func(T) K) []MapRecord[K, T] {
	result := make([]MapRecord[K, T], 0, len(seq))

	for key, val := range seq {
		result = append(result, MapRecord[K, T]{key, val})
	}

	return result
}
