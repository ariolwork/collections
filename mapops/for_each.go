package mapops

func ForEach[K comparable, V any](seq map[K]V, execFunc func(K, V)) {
	for key, val := range seq {
		execFunc(key, val)
	}
}

func ForEachWithErr[K comparable, V any](seq map[K]V, execFunc func(K, V) error) error {
	for key, val := range seq {
		err := execFunc(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
