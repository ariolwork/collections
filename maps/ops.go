package maps

import "golang.org/x/exp/constraints"

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Math Operations

func Reduce[M ~map[K]V, K comparable, V any, R any](m M, init R, cmd func(R, K, V) R) R {
	r := init

	for key, val := range m {
		r = cmd(r, key, val)
	}

	return r
}

func SumFunc[M ~map[K]V, K comparable, V any, R MathOps](m M, cmd func(K, V) R) R {
	var initVal R
	return Reduce(m, initVal, func(a R, key K, val V) R {
		return a + cmd(key, val)
	})
}

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
