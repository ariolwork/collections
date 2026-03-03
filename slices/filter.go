package slices

// The filter function select elements of a slice that satisfy to filter function.
// It creates new slice with same size and append only satisfing elements.
// If you wnaht to remove unused cpacity from slice, use slices.Clip
// Filter returns the updated slice. It is therefore necessary to store the
// result of filtering, often in the variable holding the slice itself:
//
//	slice = slices.filter(slice, filterFunc)
func Filter[S ~[]T, T any](s S, filter func(T) bool) []T {
	result := make([]T, 0, len(s))

	for _, item := range s {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}
