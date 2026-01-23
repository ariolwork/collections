package enumerate

func First[E Enumerable[T], T any](seq E) T {
	if seq.HasNext() {
		return seq.Next()
	}

	panic("Attempt to get First on empty collection. Please use FirstOrDefault")
}

func FirstOrDefault[E Enumerable[T], T any](seq E) T {
	var val T

	if seq.HasNext() {
		val = seq.Next()
	}

	return val
}

func Last[E Enumerable[T], T any](seq E) T {
	setupped := false
	var val T

	for seq.HasNext() {
		val = seq.Next()
		setupped = true
	}

	if setupped {
		return val
	}

	panic("Attempt to get Last on empty collection. Please use LastOrDefault")
}

func LastOrDefault[E Enumerable[T], T any](seq E) T {
	var val T

	for seq.HasNext() {
		val = seq.Next()
	}

	return val
}

func FirstBy[E Enumerable[T], T any](seq E, filter func(T) bool) T {
	for seq.HasNext() {
		val := seq.Next()
		if filter(val) {
			return val
		}
	}

	panic("Attempt to get FirstBy without success collection containing. Please use FirstOrDefaultBy")
}

func FirstOrDefaultBy[E Enumerable[T], T any](seq E, filter func(T) bool) T {
	var val T

	for seq.HasNext() {
		tmp := seq.Next()
		if filter(tmp) {
			val = tmp
		}
	}

	return val
}

func LastBy[E Enumerable[T], T any](seq E, filter func(T) bool) T {
	setupped := false
	var val T

	for seq.HasNext() {
		tmp := seq.Next()

		if filter(tmp) {
			val = tmp
			setupped = true
		}
	}

	if setupped {
		return val
	}

	panic("Attempt to get LastBy without success collection containing. Please use LastOrDefaultBy")
}

func LastOrDefaultBy[E Enumerable[T], T any](seq E, filter func(T) bool) T {
	var val T

	for seq.HasNext() {
		tmp := seq.Next()
		if filter(tmp) {
			val = tmp
		}
	}

	return val
}
