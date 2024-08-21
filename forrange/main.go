package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
		}()
	}
	<-time.Tick(time.Second)
}
