package maps

// The Filter function select elements of a map that satisfy to filter function.
// It creates new map with same size and append only satisfing elements.
// Filter returns the updated map. It is therefore necessary to store the
// result of filtering, often in the variable holding the map itself:
//
//	m = maps.Filter(m, filterFunc)
func Filter[M ~map[K]V, K comparable, V any](m M, filter func(K, V) bool) M {
	result := make(map[K]V, len(m))

	for key, value := range m {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}

// The DeleteFunc function select key of a map that unsatisfy to filter function.
// !It updates source map.
//
//	maps.DeleteFunc(m, filterFunc)
func DeleteFunc[M ~map[K]V, K comparable, V any](m map[K]V, filter func(K, V) bool) {
	trash := make([]K, 0, len(m))

	for key, value := range m {
		if filter(key, value) {
			trash = append(trash, key)
		}
	}

	for _, key := range trash {
		delete(m, key)
	}
}
