package sliceops

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
