package en_select

import (
	"asdanko/enumerate/enumerate"
	"iter"
)

func FromIter[T, R any](seq iter.Seq[T], applyFunc func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			if !yield(applyFunc(item)) { // yield returns false to stop early
				return
			}
		}
	}
}

func FromSlice[T, R any](seq []T, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectSlice[T, R]{subEnum: seq, applyFunc: applyFunc}
}

func From[E enumerate.Enumerable[T], T, R any](seq E, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectEnumerable[E, T, R]{subEnum: seq, applyFunc: applyFunc}
}
