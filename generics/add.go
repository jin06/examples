package main

// 加法函数，支持int类型，float类型相加
func Add[T int | float32 | float64](a T, b T) (res T) {
	return a + b
}
