package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 0, 10000)
	for i := 0; i < 10; i++ {
		go func() {
			arr = append(arr, 1)
		}()
	}

	data := map[int]string{}
	go func() {
		for i := 0; i < 100; i++ {
			for key := range data {
				fmt.Println(key)
			}
		}
	}()
	for i := 0; i < 1000000; i++ {
		data[i] = "test"
	}
}
