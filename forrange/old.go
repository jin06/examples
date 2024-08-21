package main

import "fmt"

func oldRange() {
	values := []int{1, 2, 3}
	var rec int
	for _, v := range values {
		rec = v
		go func() {
			fmt.Println(rec)
		}()
	}
}
