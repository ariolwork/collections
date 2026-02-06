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

func SliceToSlice[T, R any](seq []T, applyFunc func(T) R) []R {
	return enumerate.ToSlice(&selectSlice[T, R]{subEnum: seq, applyFunc: applyFunc})
}

func Map[M enumerate.MapRecord[K, V], K comparable, V, R any](seq map[K]V, applyFunc func(K, V) R) enumerate.Enumerable[R] {
	return &selectMap[K, V, R]{subEnum: seq, applyFunc: applyFunc}
}

func MapToSlice[M enumerate.MapRecord[K, V], K comparable, V, R any](seq map[K]V, applyFunc func(K, V) R) []R {
	return enumerate.ToSlice(&selectMap[K, V, R]{subEnum: seq, applyFunc: applyFunc})
}
