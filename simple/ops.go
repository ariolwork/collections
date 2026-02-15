package simple

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](source []T) T {
	var max T

	if len(source) == 0 {
		return max
	}

	max = source[0]

	for _, tmp := range source {
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func Min[T constraints.Ordered](source []T) T {
	var min T

	if len(source) == 0 {
		return min
	}

	min = source[0]

	for _, tmp := range source {
		if tmp < min {
			min = tmp
		}
	}

	return min
}

func MaxBy[T any](source []T, compare func(first, second T) int) T {
	var max T

	if len(source) == 0 {
		return max
	}

	max = source[0]

	for _, tmp := range source {
		if compare(tmp, max) > 0 {
			max = tmp
		}
	}

	return max
}

func MinBy[T any](source []T, compare func(first, second T) int) T {
	var min T

	if len(source) == 0 {
		return min
	}

	min = source[0]

	for _, tmp := range source {
		if compare(tmp, min) > 0 {
			min = tmp
		}
	}

	return min
}

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[T MathOps](source []T) T {
	var sum T

	for _, tmp := range source {
		sum += tmp
	}

	return sum
}

func SumBy[T any, R MathOps](source []T, selector func(T) R) R {
	var sum R

	for _, tmp := range source {
		sum += selector(tmp)
	}

	return sum
}

func Multiply[T MathOps](source []T) T {
	var mul T

	if len(source) == 0 {
		return mul
	}

	mul = source[0]

	for i := 1; i < len(source); i++ {
		mul *= source[i]
	}

	return mul
}

func MultiplyBy[T any, R MathOps](source []T, selector func(T) R) R {
	var mul R

	if len(source) == 0 {
		return mul
	}

	mul = selector(source[0])

	for i := 1; i < len(source); i++ {
		mul *= selector(source[i])
	}

	return mul
}
