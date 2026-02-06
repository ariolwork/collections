package exec

import (
	"asdanko/enumerate/enumerate"
)

func ForEach[E enumerate.Enumerable[T], T any](seq E, execFunc func(T)) {
	for seq.HasNext() {
		val := seq.Next()
		execFunc(val)
	}
}

func ForEachWithErr[E enumerate.Enumerable[T], T any](seq E, execFunc func(T) error) error {
	for seq.HasNext() {
		val := seq.Next()
		err := execFunc(val)
		if err != nil {
			return err
		}
	}

	return nil
}

func ForEachInSlice[T any](seq []T, execFunc func(T)) {
	for _, val := range seq {
		execFunc(val)
	}
}

func ForEachInSliceWithErr[T any](seq []T, execFunc func(T) error) error {
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
