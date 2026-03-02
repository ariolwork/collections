package maps

func Append[M ~map[K]V, K comparable, V any](first, second M) {
	for key, val := range second {
		first[key] = val
	}
}

func Join[M ~map[K]V, K comparable, V any](maps ...M) map[K]V {
	initSize := 0
	for _, m := range maps {
		if initSize < len(m) {
			initSize = len(m)
		}
	}

	result := make(map[K]V, initSize)

	for _, m := range maps {
		for key, val := range m {
			result[key] = val
		}
	}

	return result
}

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
