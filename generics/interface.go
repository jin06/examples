package main

type number interface {
	int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float64 | float32
}

func Add2[T number](a, b T) T {
	return a + b
}
