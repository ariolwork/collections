package slices

import (
	"golang.org/x/exp/constraints"
)

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Math Operations

// The Reduce function apply each of collection element to single value.
// Right way to use. For example to join all values in slice in raw
//
//	val = slices.Reduce([]int{1,2,3}, "", func(r string, val int) string {return r + itoa.Itoa(val)})
//
// Pay attention! It's just an example, Not recommended way to join values in string raw
func Reduce[S ~[]T, T, R any](s S, init R, cmd func(R, T) R) R {
	r := init

	for _, tmp := range s {
		r = cmd(r, tmp)
	}

	return r
}

// The Sum function sum each of collection element to single value.
func Sum[S ~[]T, T MathOps](s S) T {
	var initVal T
	return Reduce(s, initVal, func(a T, item T) T {
		return a + item
	})
}

// The SumFunc function sum function results for each collection element.
// Right way to use.
//
//	val = slices.SumFunc([]string{"aaa","bb"}, func(r string) int {return len(r)})
//
// Pay attention! It's just an example.
func SumFunc[S ~[]T, T any, R MathOps](s S, cmd func(T) R) R {
	var initVal R
	return Reduce(s, initVal, func(a R, item T) R {
		return a + cmd(item)
	})
}

// The Multiply function multiply each of collection element to single value.
func Multiply[S ~[]T, T MathOps](s S) T {
	var mul T

	if len(s) == 0 {
		return mul
	}

	mul = s[0]

	for i := 1; i < len(s); i++ {
		mul *= s[i]
	}

	return mul
}

// The Multiply function multiply function results for each collection element.
// Right way to use.
//
//	val = slices.MultiplyFunc([]string{"aaa","bb"}, func(r string) int {return len(r)})
//
// Pay attention! It's just an example.
func MultiplyFunc[S ~[]T, T any, R MathOps](s S, cmd func(T) R) R {
	var mul R

	if len(s) == 0 {
		return mul
	}

	mul = cmd(s[0])

	for i := 1; i < len(s); i++ {
		mul *= cmd(s[i])
	}

	return mul
}

// First Last Operations

// The First function returns first element of slice.
// If collection is empty, function panic
// You should use FirstOrDefault function if collection can be empty
func First[S ~[]T, T any](s S) T {
	if len(s) == 0 {
		panic("Attempt to get First on empty collection. Please use FirstOrDefault")
	}

	return s[0]
}

// The FirstOrDefault function returns first element of slice.
// If collection is empty, function returns default value
func FirstOrDefault[S ~[]T, T any](s S) T {
	var val T

	if len(s) != 0 {
		val = s[0]
	}

	return val
}

// The FirstFunc function returns first element of slice satisfying filter (cmd).
// If there is no satisfying elements in collection, function panic
// You should use FirstOrDefaultFunc function to avoid panics
func FirstFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	for _, item := range s {
		if cmd(item) {
			return item
		}
	}

	panic("Attempt to get FirstBy without success collection containing. Please use FirstOrDefaultBy")
}

// The FirstOrDefaultFunc function returns first element of slice satisfying filter (cmd).
// If there is no satisfying elements in collection, function returns default value
func FirstOrDefaultFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	for _, item := range s {
		if cmd(item) {
			return item
		}
	}

	var val T

	return val
}

// The Last function returns last element of slice.
// If collection is empty, function panic
// You should use LastOrDefault function if collection can be empty
func Last[S ~[]T, T any](s S) T {
	if len(s) == 0 {
		panic("Attempt to get Last on empty collection. Please use LastOrDefault")
	}

	return s[len(s)-1]
}

// The LastOrDefault function returns last element of slice.
// If collection is empty, function returns default value
func LastOrDefault[S ~[]T, T any](s S) T {
	var val T

	if len(s) != 0 {
		val = s[len(s)-1]
	}

	return val
}

// The LastFunc function returns last element of slice satisfying filter (cmd).
// If there is no satisfying elements in collection, function panic
// You should use FirstOrDefaultFunc function to avoid panics
func LastFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	setupped := false
	var val T

	for i := len(s) - 1; i != -1; i-- {
		tmp := s[i]

		if cmd(tmp) {
			val = tmp
			setupped = true
			break
		}
	}

	if setupped {
		return val
	}

	panic("Attempt to get LastBy without success collection containing. Please use LastOrDefaultBy")
}

// The LastOrDefaultFunc function returns last element of slice satisfying filter (cmd).
// If there is no satisfying elements in collection, function returns default value
func LastOrDefaultFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	var val T

	for i := len(s) - 1; i != -1; i-- {
		tmp := s[i]

		if cmd(tmp) {
			val = tmp
			break
		}
	}

	return val
}

// The All function accept slice with bool elements and return true if all are true.
func All[S ~[]T, T ~bool](s S) bool {
	for _, val := range s {
		if !val {
			return false
		}
	}

	return true
}

// The AllFunc function returns true if cmd function return true for all elements.
// If collection is empty, true returns. So you may need to check collection on empty.
// Right way to use
//
// if slices.AllFunc([]int{2,4,6}, func(val int)bool { return val&1 == 0 } ) { ...
func AllFunc[S ~[]T, T any](s S, cmd func(T) bool) bool {
	for _, val := range s {
		if !cmd(val) {
			return false
		}
	}

	return true
}

// The Any function accept slice with bool elements and return true if one or more are true.
func Any[S ~[]T, T ~bool](s S) bool {
	for _, val := range s {
		if val {
			return true
		}
	}

	return false
}

// The AnyFunc function returns true if cmd function return true for one or more.
// If collection is empty, false returns. So you may need to check collection on empty.
// Right way to use
//
// if slices.AnyFunc([]int{1,2,3,4}, func(val int)bool { return val&1 == 0 } ) { ...
func AnyFunc[S ~[]T, T any](s S, cmd func(T) bool) bool {
	for _, val := range s {
		if cmd(val) {
			return true
		}
	}

	return false
}

// The IsEmpty function checks if collection is empty.
// It returns true if collection is empty
func IsEmpty[S ~[]T, T any](s S) bool {
	return len(s) == 0
}

// The NotEmpty function checks if collection is empty.
// It returns true if collection is not empty
func NotEmpty[S ~[]T, T any](s S) bool {
	return len(s) != 0
}
