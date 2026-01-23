package en_select

import (
	"iter"

	"asdanko/enumerate/enumerate"
)

func FromIter[T, R any](seq iter.Seq[T], applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectIters[T, R]{subEnum: seq, applyFunc: applyFunc}
}

func FromSlice[T, R any](seq []T, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectSlice[T, R]{subEnum: seq, applyFunc: applyFunc}
}

func From[E enumerate.Enumerable[T], T, R any](seq E, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectEnumerable[E, T, R]{subEnum: seq, applyFunc: applyFunc}
}
