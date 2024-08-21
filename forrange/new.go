package main

import (
	"fmt"
	"time"
)

func newRange() {
	values := []int{1, 2, 3}
	for _, v := range values {
		var rec = v
		go func() {
			fmt.Println(rec)
		}()
	}

	<-time.Tick(time.Second)
}
