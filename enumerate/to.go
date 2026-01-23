package enumerate

func ToSlice[E Enumerable[T], T any](seq E) []T {
	result := make([]T, 0, seq.GetLen())

	for seq.HasNext() {
		result = append(result, seq.Next())
	}

	return result
}
