package main

import "testing"

func TestConstraints(t *testing.T) {
	t.Log(Add3(1.2, 2.3))
	var a, b complex64 = 1.2, 2.3
	t.Log(Add3(a, b))
}
