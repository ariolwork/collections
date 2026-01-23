package en_select

import (
	"asdanko/enumerate/enumerate"
)

func FromSlice[T, R any](seq []T, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectSlice[T, R]{subEnum: seq, applyFunc: applyFunc}
}

func From[E enumerate.Enumerable[T], T, R any](seq E, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectEnumerable[E, T, R]{subEnum: seq, applyFunc: applyFunc}
}
