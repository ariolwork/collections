package enumerate

func ToSlice[E Enumerable[T], T any](seq E) []T {
	result := make([]T, 0, seq.GetLen())

	for seq.HasNext() {
		result = append(result, seq.Next())
	}

	return result
}

func ToMap[E Enumerable[T], T any, K comparable](seq E, keySelector func(T) K) map[K]T {
	result := make(map[K]T, seq.GetLen())

	for seq.HasNext() {
		val := seq.Next()
		result[keySelector(val)] = val
	}

	return result
}

func ToMapBuckets[E Enumerable[T], T any, K comparable](seq E, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T, seq.GetLen())

	for seq.HasNext() {
		val := seq.Next()
		result[keySelector(val)] = append(result[keySelector(val)], val)
	}

	return result
}
