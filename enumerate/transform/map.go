package transform

import "asdanko/enumerate/enumerate/extract"

type selectMap[K comparable, V, R any] struct {
	currentId int
	subEnum   map[K]V
	applyFunc func(K, V) R
	keys      []K
}

func newSelectMap[K comparable, V, R any](subEnum map[K]V, applyFunc func(K, V) R) *selectMap[K, V, R] {
	return &selectMap[K, V, R]{
		subEnum:   subEnum,
		applyFunc: applyFunc,
		keys:      extract.MapKeys(subEnum),
	}
}

func (s *selectMap[K, V, R]) GetLen() int {
	return len(s.subEnum)
}

func (s *selectMap[K, V, R]) Next() R {
	if s.currentId >= len(s.subEnum) {
		panic("out of range access attempt")
	}

	key := s.keys[s.currentId]
	val := s.subEnum[key]

	s.currentId += 1

	return s.applyFunc(key, val)
}

func (s *selectMap[K, V, R]) HasNext() bool {
	return s.currentId < len(s.keys)
}
