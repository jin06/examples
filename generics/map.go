package main

import "fmt"

func printmaps[Key string | int, Val int | float64](data map[Key]Val) {
	for k, v := range data {
		fmt.Printf("name: %v age: %v \n", k, v)
	}
}
