package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Log(Add[int](1, 2))
	t.Log(Add(2, 2))
	t.Log(Add(1.1, 2.2))
}
