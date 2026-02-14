package split

import "asdanko/enumerate/enumerate"

func newBatchSlice[T any](seq []T, batchSize int) enumerate.Enumerable[[]T] {
	return &batchSlice[T]{subEnum: seq, batchSize: batchSize, totalBatches: (len(seq) + batchSize - 1) / batchSize}
}

type batchSlice[T any] struct {
	subEnum   []T
	batchSize int

	totalBatches int
	currentBatch int
}

func (s *batchSlice[T]) GetLen() int {
	return len(s.subEnum)
}

func (s *batchSlice[T]) Next() []T {
	if s.currentBatch > s.batchSize {
		panic("attempt to request batch out of range")
	}

	start := s.currentBatch * s.batchSize

	s.currentBatch += 1

	end := s.currentBatch * s.batchSize
	if end > len(s.subEnum) {
		end = len(s.subEnum)
	}

	return s.subEnum[start:end]
}

func (s *batchSlice[T]) HasNext() bool {
	return s.currentBatch < s.totalBatches
}
