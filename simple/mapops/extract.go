package mapops

func SelectValues[K comparable, V any](m map[K]V) []V {
	resultList := make([]V, 0, len(m))
	for _, v := range m {
		resultList = append(resultList, v)
	}

	return resultList
}

func SelectKeys[K comparable, V any](m map[K]V) []K {
	resultList := make([]K, 0, len(m))
	for k, _ := range m {
		resultList = append(resultList, k)
	}

	return resultList
}
