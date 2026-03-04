package maps

// The UpdateValues function convert each element value of map.
// It creates new map with updated vavlues. If current map
// values updating required, use UpdateValues
// Right way to use:
//
//	map := map[int]int{1:1, 2:2}
//	mapOfStrings := maps.TransformValues(map, func(k, v int) string {return itoa.Itoa(v)})
func TransformValues[M ~map[K]V, K comparable, V, R any](m M, applyFunc func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))

	for key, val := range m {
		result[key] = applyFunc(key, val)
	}

	return result
}

// The TransformKeys function convert each element key of map.
// It creates new map with updated vavlues.
// !!!If new keys have similar values, values will be overriten.
// To save all values, use Rotate function
// Right way to use:
//
//	map := map[int]int{1:1, 2:2}
//	stringKeysMap := maps.TransformKeys(map, func(k, v int) string {return itoa.Itoa(k)})
func TransformKeys[M ~map[K]V, K, R comparable, V any](m M, applyFunc func(K, V) R) map[R]V {
	result := make(map[R]V, len(m))

	for key, val := range m {
		result[applyFunc(key, val)] = val
	}

	return result
}

// The Rotate function select new keys for each element of map.
// It creates new rotated map.
// !!!If new keys have similar values, values will be saved in slice
// Right way to use:
//
//	map := map[int]int{1:1, 2:2}
//	stringKeysMap := maps.Rotate(map, func(k, v int) string {return itoa.Itoa(k)})
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
