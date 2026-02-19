package split

import "asdanko/enumerate/enumerate"

// todo а что если длина нижележащего не определена

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

	ind := 0
	for s.subEnum.HasNext() && ind < s.batchSize {
		batch = append(batch, s.subEnum.Next())
		ind += 1
	}

	return batch
}

func (s *batchEnumerable[E, T]) HasNext() bool {
	return s.currentBatch < s.totalBatches
}
