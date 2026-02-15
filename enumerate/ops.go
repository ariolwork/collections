package enumerate

import "golang.org/x/exp/constraints"

func Max[E Enumerable[T], T constraints.Ordered](seq E) T {
	setted := false
	var max T

	for seq.HasNext() {
		tmp := seq.Next()
		if setted {
			if tmp > max {
				max = tmp
			}
		} else {
			max = tmp
			setted = true
		}
	}

	return max
}

func Min[E Enumerable[T], T constraints.Ordered](seq E) T {
	setted := false
	var min T

	for seq.HasNext() {
		tmp := seq.Next()
		if setted {
			if tmp < min {
				min = tmp
			}
		} else {
			min = tmp
			setted = true
		}
	}

	return min
}

func MaxBy[E Enumerable[T], T any](seq E, compare func(first, second T) int) T {
	setted := false
	var max T

	for seq.HasNext() {
		tmp := seq.Next()
		if setted {
			if compare(tmp, max) > 0 {
				max = tmp
			}
		} else {
			max = tmp
			setted = true
		}
	}

	return max
}

func MinBy[E Enumerable[T], T any](seq E, compare func(first, second T) int) T {
	setted := false
	var min T

	for seq.HasNext() {
		tmp := seq.Next()
		if setted {
			if compare(tmp, min) < 0 {
				min = tmp
			}
		} else {
			min = tmp
			setted = true
		}
	}

	return min
}

type MathOps interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[E Enumerable[T], T MathOps](seq E) T {
	var sum T

	for seq.HasNext() {
		sum += seq.Next()
	}

	return sum
}

func SumBy[E Enumerable[T], T any, R MathOps](seq E, selector func(T) R) R {
	var sum R

	for seq.HasNext() {
		sum += selector(seq.Next())
	}

	return sum
}

// todo это By да и преализация неверна, просуммируем  0
func Multiply[E Enumerable[T], T any, R MathOps](seq E, selector func(T) R) R {
	var sum R

	for seq.HasNext() {
		sum *= selector(seq.Next())
	}

	return sum
}
