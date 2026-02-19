package mapops

//todo переименовать все seq в source
//todo ротировать функции и обекты над которыми приемняем для соответствися лучшим

func TransformValues[K comparable, V, R any](source map[K]V, applyFunc func(K, V) R) map[K]R {
	result := make(map[K]R, len(source))

	for key, val := range source {
		result[key] = applyFunc(key, val)
	}

	return result
}
