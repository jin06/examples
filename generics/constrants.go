package main

import "golang.org/x/exp/constraints"

func Add3[T constraints.Integer | constraints.Complex | constraints.Float](a T, b T) T {
	return a + b
}
