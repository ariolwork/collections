package transform

import (
	"asdanko/enumerate/enumerate"
)

func Enumerable[E enumerate.Enumerable[T], T, R any](seq E, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectEnumerable[E, T, R]{subEnum: seq, applyFunc: applyFunc}
}

func Slice[T, R any](seq []T, applyFunc func(T) R) enumerate.Enumerable[R] {
	return &selectSlice[T, R]{subEnum: seq, applyFunc: applyFunc}
}

func Map[K comparable, V, R any](seq map[K]V, applyFunc func(K, V) R) enumerate.Enumerable[R] {
	return &selectMap[K, V, R]{subEnum: seq, applyFunc: applyFunc}
}
