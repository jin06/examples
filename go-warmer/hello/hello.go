package main

import (
	"fmt"
)

//export Add
func Add(a int32, b int32) int32 {
	fmt.Println("Hello from Go!")
	return a + b
}

//export HelloWorld
func HelloWorld() {
	fmt.Println("Hello, World!")
}

func main() {
}
