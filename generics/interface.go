package main

type addInterface interface {
	int | string | float64
}

func Add2[T addInterface](a, b T) T {
	return a + b
}
