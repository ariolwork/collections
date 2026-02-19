package maps

//todo переименовать все seq в source
//todo ротировать функции и обекты над которыми приемняем для соответствися лучшим

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
