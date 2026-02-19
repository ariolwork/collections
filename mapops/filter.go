package mapops

func Filter[K comparable, V any](source map[K]V, filter func(K, V) bool) map[K]V {
	result := make(map[K]V, len(source))

	for key, value := range source {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}
