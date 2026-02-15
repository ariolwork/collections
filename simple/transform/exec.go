package transform

//todo переименовать все seq в source
//todo ротировать функции и обекты над которыми приемняем для соответствися лучшим прктикам

func Slice[T, R any](source []T, applyFunc func(T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, applyFunc(item))
	}

	return result
}

func MapValues[K comparable, V, R any](source map[K]V, applyFunc func(K, V) R) map[K]R {
	result := make(map[K]R, len(source))

	for key, val := range source {
		result[key] = applyFunc(key, val)
	}

	return result
}
