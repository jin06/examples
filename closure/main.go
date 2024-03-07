package main

import (
	"fmt"
)

func NewCounter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	count := NewCounter()
	for i := 0; i < 10; i++ {
		fmt.Println(count())
	}
}
