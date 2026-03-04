package maps

// The ForEach function exec any func for each element of map.
// !!!Pay attention, if you want to change map non ref values you should use TransformValues, UpdateValues or TransformKeys
//
// Right way to use this function
//
//	s := map[int]int{1:1, 2:2}
//	maps.ForEach(s, func(k, v int){ logger.Log("Map record", k, v) })
func ForEach[M ~map[K]V, K comparable, V any](m M, cmd func(K, V)) {
	for key, val := range m {
		cmd(key, val)
	}
}

// The ForEach function exec any func for each element of map unitl error returns.
// !!!Pay attention, if you want to change map non ref values you should use TransformValues, UpdateValues or TransformKeys
//
// Right way to use this function
//
//	s := map[int]int{1:1, 2:2}
//	maps.ForEach(s, func(k, v int) error {
//		if k == 0 {
//			return errors.New("Unsupported value")
//		}
//		logger.Log("Map record", k, v)
//		return nil
//	})
func ForEachWithErr[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) error) error {
	for key, val := range m {
		err := cmd(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
