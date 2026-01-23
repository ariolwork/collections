package en_select

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"asdanko/enumerate/enumerate"

	"github.com/stretchr/testify/assert"
)

func TestSelectForSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	r := enumerate.ToSlice(FromSlice(items, func(i int) string {
		return fmt.Sprintf("A: %d", i)
	}))

	assert.ElementsMatch(t, r, []string{"A: 1", "A: 2", "A: 3", "A: 4", "A: 5"})
}

func TestSelectMax(t *testing.T) {
	maxVal := time.Now().Add(time.Hour * 10)
	items := []time.Time{time.Now().Add(time.Hour * 1), maxVal, time.Now().Add(time.Hour * 2), time.Now().Add(time.Hour * 3)}

	r := enumerate.MaxBy(
		FromSlice(items, func(i time.Time) time.Time {
			return i.Add(time.Duration(time.Minute))
		}),
		func(a, b time.Time) int {
			return a.Compare(b)
		},
	)

	assert.Equal(t, maxVal.Add(time.Minute), r)
}

func BenchmarkSelectForSliceTest(b *testing.B) {
	rnd := rand.New(rand.NewSource(int64(100)))
	len := rnd.Intn(1000)
	source := make([]int, 0, len)
	ex := make([]int, 0, len)

	for i := 0; i < len; i++ {
		source = append(source, i)
		ex = append(ex, i+3)
	}

	for i := 0; i < b.N; i++ {

		r := enumerate.ToSlice(
			From(
				From(
					FromSlice(
						source,
						func(i int) int {
							return i + 1
						},
					),
					func(i int) int {
						return i + 1
					},
				),
				func(i int) int {
					return i + 1
				},
			),
		)
		_ = r
		//assert.ElementsMatch(b, r, ex)
	}
}
