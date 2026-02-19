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

func TransformKeys[K, R comparable, V any](source map[K]V, applyFunc func(K, V) R) map[R]V {
	result := make(map[R]V, len(source))

	for key, val := range source {
		result[applyFunc(key, val)] = val
	}

	return result
}

func Rotate[K, R comparable, V any](source map[K]V, applyFunc func(K, V) R) map[R][]V {
	result := make(map[R][]V, len(source))

	for key, val := range source {
		result[applyFunc(key, val)] = append(result[applyFunc(key, val)], val)
	}

	return result
}
