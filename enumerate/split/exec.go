package split

import (
	"asdanko/enumerate/enumerate"
)

func ByBatches[E enumerate.Enumerable[T], T any](seq E, size int) enumerate.Enumerable[[]T] {
	return newBatchEnumerable(seq, size)
}

func ByBatchesSlice[T any](seq []T, size int) enumerate.Enumerable[[]T] {
	return newBatchSlice(seq, size)
}
