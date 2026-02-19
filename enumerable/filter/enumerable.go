package filter

import "asdanko/enumerate/enumerate"

func newEnumerable[E enumerate.Enumerable[T], T any](seq E, f func(T) bool) enumerate.Enumerable[T] {
	return &filteredEnumerable[E, T]{subEnum: seq}
}

type filteredEnumerable[E enumerate.Enumerable[T], T any] struct {
	subEnum E
	f       func(T) bool

	nextItem       T
	nextItemFilled bool
}

func (s *filteredEnumerable[E, T]) GetLen() int {
	// size become unknown
	return 0
}

func (s *filteredEnumerable[E, T]) Next() T {
	if !s.nextItemFilled {

		item, found := s.getNext()
		if !found {
			panic("Attempt to get next item, nothing found by filter")
		}

		return item
	}

	toReturn := s.nextItem
	s.nextItem, s.nextItemFilled = s.getNext()

	return toReturn
}

func (s *filteredEnumerable[E, T]) HasNext() bool {
	if s.nextItemFilled {
		return true
	}

	s.nextItem, s.nextItemFilled = s.getNext()

	return s.nextItemFilled

}

func (s *filteredEnumerable[E, T]) getNext() (next T, found bool) {
	for s.subEnum.HasNext() {
		next = s.subEnum.Next()
		if s.f(next) {
			found = true
			return
		}
	}

	found = false
	return
}
