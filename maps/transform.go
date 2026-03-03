package maps

func TransformValues[M ~map[K]V, K comparable, V, R any](m M, applyFunc func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))

	for key, val := range m {
		result[key] = applyFunc(key, val)
	}

	return result
}

func TransformKeys[M ~map[K]V, K, R comparable, V any](m M, applyFunc func(K, V) R) map[R]V {
	result := make(map[R]V, len(m))

	for key, val := range m {
		result[applyFunc(key, val)] = val
	}

	return result
}

func Rotate[M ~map[K]V, K, R comparable, V any](m M, applyFunc func(K, V) R) map[R][]V {
	result := make(map[R][]V, len(m))

	for key, val := range m {
		result[applyFunc(key, val)] = append(result[applyFunc(key, val)], val)
	}

	return result
}

// The UpdateValues function convert each element of map.
// It update current map values
// Right way to use:
//
//	map := map[int]int{1:1, 2:2}
//	maps.UpdateValues(map, func(k, v int) int {return v * 2})
func UpdateValues[M ~map[K]V, K comparable, V any](m M, applyFunc func(K, V) V) {
	for key, val := range m {
		m[key] = applyFunc(key, val)
	}
}
