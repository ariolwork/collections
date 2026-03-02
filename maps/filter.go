package maps

func Filter[M ~map[K]V, K comparable, V any](m M, filter func(K, V) bool) M {
	result := make(map[K]V, len(m))

	for key, value := range m {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}

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
