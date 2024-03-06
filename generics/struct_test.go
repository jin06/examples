package main

import "testing"

func TestQueue(t *testing.T) {
	q := Queue[int]{Data: []int{}}
	q.Push(1, 2, 3, 4)
	q.Pop(2)
	t.Log(q.Data)
}
