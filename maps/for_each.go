package maps

func ForEach[M ~map[K]V, K comparable, V any](m M, cmd func(K, V)) {
	for key, val := range m {
		cmd(key, val)
	}
}

func ForEachWithErr[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) error) error {
	for key, val := range m {
		err := cmd(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
