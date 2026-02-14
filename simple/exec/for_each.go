package exec

func ForEach[T any](seq []T, execFunc func(T)) {
	for _, val := range seq {
		execFunc(val)
	}
}

func ForEachWithErr[T any](seq []T, execFunc func(T) error) error {
	for _, val := range seq {
		err := execFunc(val)
		if err != nil {
			return err
		}
	}

	return nil
}

func ForEachInMap[K comparable, V any](seq map[K]V, execFunc func(K, V)) {
	for key, val := range seq {
		execFunc(key, val)
	}
}

func ForEachInMapWithErr[K comparable, V any](seq map[K]V, execFunc func(K, V) error) error {
	for key, val := range seq {
		err := execFunc(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
