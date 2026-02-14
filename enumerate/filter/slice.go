package filter

import "asdanko/enumerate/enumerate"

func newSlice[T any](seq []T, f func(T) bool) enumerate.Enumerable[T] {
	return &filteredSlice[T]{subEnum: seq}
}

type filteredSlice[T any] struct {
	subEnum []T
	f       func(T) bool

	nextItem       T
	nextItemFilled bool
	currentInd     int
}

func (s *filteredSlice[T]) GetLen() int {
	// size become unknown
	return 0
}

func (s *filteredSlice[T]) Next() T {
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

func (s *filteredSlice[T]) HasNext() bool {
	if s.nextItemFilled {
		return true
	}

	s.nextItem, s.nextItemFilled = s.getNext()

	return s.nextItemFilled

}

func (s *filteredSlice[T]) getNext() (next T, found bool) {
	for s.currentInd < len(s.subEnum) {
		next = s.subEnum[s.currentInd]
		s.currentInd += 1
		if s.f(next) {
			found = true
			return
		}
	}

	found = false
	return
}
