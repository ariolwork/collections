package en_select

import (
	"asdanko/enumerate/enumerate"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectForIter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	r := enumerate.ToSlice(FromSlice(items, func(i int) string {
		return fmt.Sprintf("A: %d", i)
	}))

	assert.ElementsMatch(t, r, []string{"A: 1", "A: 2", "A: 3", "A: 4", "A: 5"})
}
