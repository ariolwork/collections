package filter

import (
	"asdanko/enumerate/enumerate"
)

func By[E enumerate.Enumerable[T], T any](seq E, f func(T) bool) enumerate.Enumerable[T] {
	return newEnumerable(seq, f)
}

func SliceBy[T any](seq []T, f func(T) bool) enumerate.Enumerable[T] {
	return newSlice(seq, f)
}
