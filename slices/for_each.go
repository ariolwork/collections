package slices

// The ForEach function exec any func for each element of slice.
// !!!Pay attention, if you want to change slice non ref values you should use Transform or Update
// This example won't update source slice, just take you memory and cpu time:
//
//	s := []int{1,2,3}
//	slices.ForEach(s, func(i int){ i+=2 })
//
// Right way to use this function
//
//	s := []int{1,2,3}
//	slices.ForEach(s, func(i int){ logger.Log("Element", i) })
func ForEach[S ~[]T, T any](s S, cmd func(T)) {
	for _, val := range s {
		cmd(val)
	}
}

// The ForEachWithErr close to ForEach function.
// It exec any func for each element of slice unitl error returns.
// Right way to use this function
//
//	s := []int{1,2,3}
//	slices.ForEach(s, func(i int) error {
//		if i == 5{
//			return "Unexpected value"
//		}
//		logger.Log("Element", i)
//		return nil
//	})
func ForEachWithErr[S ~[]T, T any](s []T, cmd func(T) error) error {
	for _, val := range s {
		err := cmd(val)
		if err != nil {
			return err
		}
	}

	return nil
}
