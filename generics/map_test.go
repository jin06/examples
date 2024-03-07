package main

import "testing"

func TestPrint(t *testing.T) {
	data := map[string]int{"xiaowang": 21, "xiaoli": 32}
	printmaps(data)
	data2 := map[string]float64{"laozhang": 55.0, "laoliu": 90.0}
	printmaps(data2)
}
