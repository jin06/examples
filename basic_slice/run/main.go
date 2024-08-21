package main

import "fmt"

func main() {

	// values := []int{}

	// values := make([]int, 0, 3)
	values := make([]int, 3, 3)

	values2 := values[1:2]
	fmt.Println(values2)
}
