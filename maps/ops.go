package maps

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Math Operations

// The Reduce function apply each of collection element to single value.
// Right way to use. For example to join all values in map in raw
//
//	val = maps.Reduce(map[int]int{1:1, 2:2}, "", func(r string, k, v int) string {return r + fmt.Sprintf("key: %d, val: %d; ", k, v)})
//
// Pay attention! It's just an example, Not recommended way to join values in string raw
func Reduce[M ~map[K]V, K comparable, V any, R any](m M, init R, cmd func(R, K, V) R) R {
	r := init

	for key, val := range m {
		r = cmd(r, key, val)
	}

	return r
}

// The SumFunc function sum function results for each collection element.
// Right way to use.
//
//	val = maps.SumFunc(map[int]int{1:1, 2:2}, "", func(r string, k, v int) int {return k*v})
//
// Pay attention! It's just an example.
func SumFunc[M ~map[K]V, K comparable, V any, R MathOps](m M, cmd func(K, V) R) R {
	var initVal R
	return Reduce(m, initVal, func(a R, key K, val V) R {
		return a + cmd(key, val)
	})
}

// The MultiplyFunc function multiply function results for each collection element.
// Right way to use.
//
//	val = maps.MultiplyFunc(map[int]int{1:1, 2:2}, "", func(r string, k, v int) int {return k+v})
//
// Pay attention! It's just an example.
func MultiplyFunc[M ~map[K]V, K comparable, V any, R MathOps](m M, cmd func(K, V) R) R {
	var mul R
	isDefined := false

	for key, value := range m {
		if isDefined {
			mul *= cmd(key, value)
		} else {
			mul = cmd(key, value)
		}
	}

	return mul
}

// The First function returns first element of map.
// If collection is empty, function panic. So you may need to check if collection is empty.
func First[M ~map[K]V, K comparable, V any](m M) (K, V) {
	for key, value := range m {
		return key, value
	}

	panic("Attempt to get First on empty collection. Please use FirstOrDefault")
}

// The FirstFunc function returns first element of map satisfying filter (cmd).
// If there is no satisfying elements in collection, function panic
// You should use FirstOrErrFunc function to avoid panics
func FirstFunc[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) bool) (K, V) {
	for key, value := range m {
		if cmd(key, value) {
			return key, value
		}
	}

	panic("Attempt to get FirstBy without success collection containing. Please use FirstOrDefaultBy")
}

// The FirstFunc function returns first element of map satisfying filter (cmd).
// If there is no satisfying elements in collection, function panic
func FirstOrErrFunc[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) bool) (k K, v V, err error) {
	for key, value := range m {
		if cmd(key, value) {
			return key, value, nil
		}
	}

	err = errors.New("There are no satisfying elements in collection")
	return
}

// The AllFunc function returns true if cmd function return true for all elements.
// If collection is empty, true returns. So you may need to check if collection is empty.
// Right way to use
//
// if maps.AllFunc(map[int]struct{}{}, func(key int, val struct{}) bool { return key&1 == 0 } ) { ...
func AllFunc[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) bool) bool {
	for key, val := range m {
		if !cmd(key, val) {
			return false
		}
	}

	return true
}

// The AnyFunc function returns true if cmd function return true for one or more.
// If collection is empty, false returns. So you may need to check collection on empty.
// Right way to use
//
// if maps.AnyFunc(map[int]struct{}{}, func(key int, val struct{}) bool { return key&1 == 0 } ) { ...
func AnyFunc[M ~map[K]V, K comparable, V any](m M, cmd func(K, V) bool) bool {
	for key, val := range m {
		if cmd(key, val) {
			return true
		}
	}

	return false
}

// The IsEmpty function checks if collection is empty.
// It returns true if collection is empty
func IsEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) == 0
}

// The NotEmpty function checks if collection is empty.
// It returns true if collection is not empty
func NotEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) != 0
}
