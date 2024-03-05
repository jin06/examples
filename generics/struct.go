package main

type Queue[T int | float32 | uint] struct {
	Data []T
}

func (q *Queue[T]) Push(values ...T) {
	q.Data = append(q.Data, values...)
}

func (q *Queue[T]) Pop(n int) {
	q.Data = q.Data[:len(q.Data)-n]
}
