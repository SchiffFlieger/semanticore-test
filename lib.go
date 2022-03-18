package lib

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[T Number](nums ...T) T {
	var sum T
	for _, v := range nums {
		sum += v
	}
	return sum
}
