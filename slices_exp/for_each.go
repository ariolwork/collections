package slices_exp

func ForEach[S ~[]T, T any](s S, cmd func(T)) {
	for _, val := range s {
		cmd(val)
	}
}

func ForEachWithErr[S ~[]T, T any](s []T, cmd func(T) error) error {
	for _, val := range s {
		err := cmd(val)
		if err != nil {
			return err
		}
	}

	return nil
}
