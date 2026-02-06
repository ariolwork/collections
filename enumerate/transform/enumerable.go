package transform

import "asdanko/enumerate/enumerate"

type selectEnumerable[E enumerate.Enumerable[T], T, R any] struct {
	subEnum   E
	applyFunc func(T) R
}

func (s *selectEnumerable[E, T, R]) GetLen() int {
	return s.subEnum.GetLen()
}

func (s *selectEnumerable[E, T, R]) Next() R {
	return s.applyFunc(s.subEnum.Next())
}

func (s *selectEnumerable[E, T, R]) HasNext() bool {
	return s.subEnum.HasNext()
}
