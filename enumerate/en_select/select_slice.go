package en_select

type selectSlice[T, R any] struct {
	currentId int
	subEnum   []T
	applyFunc func(T) R
}

func (s *selectSlice[T, R]) GetLen() int {
	return len(s.subEnum)
}

func (s *selectSlice[T, R]) Next() R {
	if s.currentId >= len(s.subEnum) {
		panic("out of range access attempt")
	}

	val := s.subEnum[s.currentId]

	s.currentId += 1

	return s.applyFunc(val)
}

func (s *selectSlice[T, R]) HasNext() bool {
	return s.currentId < len(s.subEnum)
}
