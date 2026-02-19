package maps

import (
	"iter"
)

func SelectValues[M ~map[K]V, K comparable, V any](m M) []V {
	resultList := make([]V, 0, len(m))
	for _, v := range m {
		resultList = append(resultList, v)
	}

	return resultList
}

func SelectKeys[M ~map[K]V, K comparable, V any](m M) []K {
	resultList := make([]K, 0, len(m))
	for k, _ := range m {
		resultList = append(resultList, k)
	}

	return resultList
}

func All[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i, v := range m {
			if !yield(i, v) {
				return
			}
		}
	}
}

func Values[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

func Keys[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range m {
			if !yield(k) {
				return
			}
		}
	}
}
