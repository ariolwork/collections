package split

import "asdanko/enumerate/enumerate"

func newBatchEnumerable[E enumerate.Enumerable[T], T any](seq E, batchSize int) enumerate.Enumerable[[]T] {
	return &batchEnumerable[E, T]{subEnum: seq, batchSize: batchSize, totalBatches: (seq.GetLen() + batchSize - 1) / batchSize}
}

type batchEnumerable[E enumerate.Enumerable[T], T any] struct {
	subEnum   E
	batchSize int

	totalBatches int
	currentBatch int
}

func (s *batchEnumerable[E, T]) GetLen() int {
	return s.subEnum.GetLen()
}

func (s *batchEnumerable[E, T]) Next() []T {
	if s.currentBatch > s.batchSize {
		panic("attempt to request batch out of range")
	}

	s.currentBatch += 1

	batch := make([]T, 0, s.batchSize)

	for s.subEnum.HasNext() {
		batch = append(batch, s.subEnum.Next())
	}

	return batch
}

func (s *batchEnumerable[E, T]) HasNext() bool {
	return s.currentBatch < s.totalBatches
}
