package maps

// The Append function append values from second map to first.
// Source map changing
func Append[M ~map[K]V, K comparable, V any](first, second M) {
	for key, val := range second {
		first[key] = val
	}
}

// The Join function join values from all incoming maps to one.
// Source maps r not changing
// Equals keys override value in result map. If all values must bee
// saved, BucketingJoin should be used
// Rigth way to use
//
//	result := maps.Join(map[int]int{1:1}, map[int]int{2:1}, map[int]int{3:1})
func Join[M ~map[K]V, K comparable, V any](maps ...M) map[K]V {
	initSize := 0
	for _, m := range maps {
		initSize += len(m)
	}

	result := make(map[K]V, initSize)

	for _, m := range maps {
		for key, val := range m {
			result[key] = val
		}
	}

	return result
}

// The BucketingJoin function join values from all incoming maps to one.
// Source maps r not changing
// Equals keys put values to slice
// Rigth way to use
//
//	result := maps.BucketingJoin(map[int]int{1:1}, map[int]int{1:3, 2:1}, map[int]int{1:4, 3:1})

func BucketingJoin[M ~map[K]V, K comparable, V any](maps ...M) map[K][]V {
	initSize := 0
	for _, m := range maps {
		if initSize < len(m) {
			initSize = len(m)
		}
	}

	result := make(map[K][]V, initSize)

	for _, m := range maps {
		for key, val := range m {
			result[key] = append(result[key], val)
		}
	}

	return result
}

// The IntersectKeys function intersect keys of maps.
// Function returns keys, contaning in all of maps
func IntersectKeys[M ~map[K]V, K comparable, V any](maps ...M) []K {
	initSize := 0
	for _, m := range maps {
		if initSize < len(m) {
			initSize = len(m)
		}
	}

	countOfHits := make(map[K]int, initSize)

	for _, m := range maps {
		for key, _ := range m {
			countOfHits[key] += 1
		}
	}

	result := make([]K, 0, initSize)
	for key, count := range countOfHits {
		if count == len(maps) {
			result = append(result, key)
		}
	}

	return result
}
