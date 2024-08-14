package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3}
	for _, v := range values {
		go func() {
			fmt.Println(v)
		}()
	}
	<-time.Tick(time.Second)
	oldRange()
	<-time.Tick(time.Second)
	newRange()
	<-time.Tick(time.Second)
}

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

func newRange() {
	values := []int{1, 2, 3}
	for _, v := range values {
		var rec = v
		go func() {
			fmt.Println(rec)
		}()
	}
}
