package slices

// The ToMap function convert slice to map with keys calculated by keySelector function
// It creates new map with same size as slice. keySelector func must returns comparable type.
// If slice contain several element with equal calculated keys, value will be overriden.
// You can use ToMapBuckets to avoid it.
// Rigth way to use function
//
//	slice = slices.ToMap([]int{1,2,3}, func(val int) string { return itoa.Itoa(val) })
func ToMap[S ~[]T, T any, K comparable](seq S, keySelector func(T) K) map[K]T {
	result := make(map[K]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = val
	}

	return result
}

// The ToMapBuckets function convert slice to map with keys calculated by keySelector function
// It creates new map with same size as slice. keySelector func must returns comparable type.
// If slice contain several element with equal calculated keys, all values will be saved in slice.
// So, it can be used as Group function.
// Rigth way to use function
//
//	//First example
//	slice = slices.ToMapBuckets([]int{1,2,3,2,1}, func(val int) string { return itoa.Itoa(val) })
//
//	//Second example
//	events := []struct{d date, name string}{ ... }
//	eventsCalendar = slices.ToMapBuckets(events, func(val struct{d date, name string}) date { return val.date })
func ToMapBuckets[S ~[]T, T any, K comparable](seq S, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T, len(seq))

	for _, val := range seq {
		result[keySelector(val)] = append(result[keySelector(val)], val)
	}

	return result
}

// The Transform function convert each element of slice to new one.
// It creates new slice with same size and append all transformed values.
// Transform returns the updated slice. It is therefore necessary to store the
// result of filtering, often in the variable holding the slice itself:
//
//	slice := []int{1,2,3,4}
//	slice = slices.Transform(slice, filterFunc)
//
// It also can return slice with another type
//
//	slice := []int{1,2,3,4}
//	slice1 := slices.Transform(slice, func(val int) string {return itoa.Itoa(val)})
func Transform[S ~[]T, T, R any](source S, applyFunc func(T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, applyFunc(item))
	}

	return result
}

// The Update function convert each element of slice.
// It update current slice values
// Right way to use:
//
//	slice := []int{1,2,3,4}
//	slices.Update(slice, func(val int) int {return val * 2})
func Update[S ~[]T, T any](source S, applyFunc func(T) T) {
	for i := 0; i < len(source); i++ {
		source[i] = applyFunc(source[i])
	}

}
