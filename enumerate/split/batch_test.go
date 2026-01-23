package split

import (
	"testing"

	"asdanko/enumerate/enumerate"
	"asdanko/enumerate/enumerate/en_select"

	"github.com/stretchr/testify/assert"
)

func TestSelectForSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	r := enumerate.ToSlice(ByBatchesSlice(items, 2))

	assert.ElementsMatch(t, r, [][]int{{1, 2}, {3, 4}, {5}})

	r1 := enumerate.ToSlice(ByBatchesSlice(items, 3))

	assert.ElementsMatch(t, r1, [][]int{{1, 2, 3}, {4, 5}})
}

func TestSelectForEnumerate(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	r := enumerate.ToSlice(ByBatches(en_select.FromSlice(items, func(i int) int { return i }), 2))
	assert.ElementsMatch(t, r, [][]int{{1, 2}, {3, 4}, {5}})

	r1 := enumerate.ToSlice(ByBatches(en_select.FromSlice(items, func(i int) int { return i }), 3))

	assert.ElementsMatch(t, r1, [][]int{{1, 2, 3}, {4, 5}})
}
