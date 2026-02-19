package sliceops

// todo добавить опции
func First[T any](source []T) T {
	if len(source) == 0 {
		panic("Attempt to get First on empty collection. Please use FirstOrDefault")
	}

	return source[0]
}

func FirstOrDefault[T any](source []T) T {
	var val T

	if len(source) != 0 {
		val = source[0]
	}

	return val
}

func FirstWhere[T any](source []T, filter func(T) bool) T {
	for _, item := range source {
		if filter(item) {
			return item
		}
	}

	panic("Attempt to get FirstBy without success collection containing. Please use FirstOrDefaultBy")
}

func FirstOrDefaultWhere[T any](source []T, filter func(T) bool) T {
	for _, item := range source {
		if filter(item) {
			return item
		}
	}

	var val T

	return val
}

func Last[T any](source []T) T {
	if len(source) == 0 {
		panic("Attempt to get Last on empty collection. Please use LastOrDefault")
	}

	return source[len(source)-1]
}

func LastOrDefault[T any](source []T) T {
	var val T

	if len(source) != 0 {
		val = source[len(source)-1]
	}

	return val
}

func LastWhere[T any](source []T, filter func(T) bool) T {
	setupped := false
	var val T

	for i := len(source) - 1; i != -1; i-- {
		tmp := source[i]

		if filter(tmp) {
			val = tmp
			setupped = true
			break
		}
	}

	if setupped {
		return val
	}

	panic("Attempt to get LastBy without success collection containing. Please use LastOrDefaultBy")
}

func LastOrDefaultWhere[T any](source []T, filter func(T) bool) T {
	var val T

	for i := len(source) - 1; i != -1; i-- {
		tmp := source[i]

		if filter(tmp) {
			val = tmp
			break
		}
	}

	return val
}
