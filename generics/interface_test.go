package main

import "testing"

func TestInterface(t *testing.T) {
	t.Log(Add2(1, 2))
	t.Logf("%+v \n", Add2(float64(1.0), float64(2.5)))
}
