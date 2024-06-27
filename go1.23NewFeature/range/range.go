package main

import (
	"fmt"
)

// for k,v := range functionWithReturnsIsTwo {
//
// }
func functionWithReturnsIsTwo(yield func(index int, arr string) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i, fmt.Sprintf("RangeTwo: I'm %d ", i)) {
			return
		}
	}
}

// for k := range functionWithReturnsIsOne {
//
// }
func functionWithReturnsIsOne(yield func(index int, arr string) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i, fmt.Sprintf("RangeOne: I'm %d ", i)) {
			return
		}
	}
}

// for range functionWithReturnsIsZero {
//
// }
func functionWithReturnsIsZero(yield func(index int, arr string) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i, fmt.Sprintf("I'm %d ", i)) {
			return
		}
	}
}

func main() {
	// 1. Basic usage, accepts iterator functions:
	// func(func() bool)
	// func(func(K) bool)
	// func(func(K, V) bool)
	fmt.Println("Test basic usage")
	for i, k := range functionWithReturnsIsTwo {
		fmt.Println(i, k)
	}
	fmt.Println("Test employee example")
	eTeckDept.GoodEmployees()
	eFinaceDept.GoodEmployees()
}
