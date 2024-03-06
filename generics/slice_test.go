package main

import "testing"

func TestDouble(t *testing.T) {
	t.Log(Double([]int{1, 2, 3, 4}))
}

func TestIndex(t *testing.T) {
	t.Log(Index([]int{3, 41, 5, 1}, 5))
}
