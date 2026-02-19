package sliceops

func Transform[T, R any](source []T, applyFunc func(T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, applyFunc(item))
	}

	return result
}
