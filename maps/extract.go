package maps

import (
	"iter"
)

// Key/Values extractions

// The Values function returns slice with all values from map.
// Function returns new slice with length equals to map length.
// Right way to use.
//
//	vals = maps.Values(mapCollection)
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	resultList := make([]V, 0, len(m))
	for _, v := range m {
		resultList = append(resultList, v)
	}

	return resultList
}

// The Keys function returns slice with all keys from map.
// Function returns new slice with length equals to map length.
// Right way to use.
//
//	vals = maps.Keys(mapCollection)
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	resultList := make([]K, 0, len(m))
	for k, _ := range m {
		resultList = append(resultList, k)
	}

	return resultList
}

// The All function returns iter.Seq2 with map elemts.
// Right way to use.
//
//	iter = maps.All(mapCollection)
func All[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i, v := range m {
			if !yield(i, v) {
				return
			}
		}
	}
}

// The ValuesSeq function returns iter.Seq with map values.
// Right way to use.
//
//	iter = maps.ValuesSeq(mapCollection)
//	// for _, val := ramge iter {...}
func ValuesSeq[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// The KeysSeq function returns iter.Seq with map keys.
// Right way to use.
//
//	iter = maps.KeysSeq(mapCollection)
//	// for _, key := ramge iter {...}
func KeysSeq[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range m {
			if !yield(k) {
				return
			}
		}
	}
}
