package slices

import "golang.org/x/exp/constraints"

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Math Operations

func Reduce[S ~[]T, T, R any](s S, init R, cmd func(R, T) R) R {
	r := init

	for _, tmp := range s {
		r = cmd(r, tmp)
	}

	return r
}

func Sum[S ~[]T, T MathOps](s S) T {
	var initVal T
	return Reduce(s, initVal, func(a T, item T) T {
		return a + item
	})
}

func SumFunc[S ~[]T, T any, R MathOps](s S, cmd func(T) R) R {
	var initVal R
	return Reduce(s, initVal, func(a R, item T) R {
		return a + cmd(item)
	})
}

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

// todo добавить опции
func First[S ~[]T, T any](s S) T {
	if len(s) == 0 {
		panic("Attempt to get First on empty collection. Please use FirstOrDefault")
	}

	return s[0]
}

func FirstOrDefault[S ~[]T, T any](s S) T {
	var val T

	if len(s) != 0 {
		val = s[0]
	}

	return val
}

func FirstFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	for _, item := range s {
		if cmd(item) {
			return item
		}
	}

	panic("Attempt to get FirstBy without success collection containing. Please use FirstOrDefaultBy")
}

func FirstOrDefaultFunc[S ~[]T, T any](s S, cmd func(T) bool) T {
	for _, item := range s {
		if cmd(item) {
			return item
		}
	}

	var val T

	return val
}

func Last[S ~[]T, T any](s S) T {
	if len(s) == 0 {
		panic("Attempt to get Last on empty collection. Please use LastOrDefault")
	}

	return s[len(s)-1]
}

func LastOrDefault[S ~[]T, T any](s S) T {
	var val T

	if len(s) != 0 {
		val = s[len(s)-1]
	}

	return val
}

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
