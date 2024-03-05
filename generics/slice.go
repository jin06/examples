package main

func Double[T any](data []T) []T {
	result := make([]T, 2*len(data))

	for i := 0; i < len(data); i++ {
		result[2*i], result[2*i+1] = data[i], data[i]
	}
	return result
}

func Index[T comparable](data []T, value T) int {
	for i := 0; i < len(data); i++ {
		if data[i] == value {
			return i
		}
	}
	return -1
}
