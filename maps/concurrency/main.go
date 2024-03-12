package main

import (
	"fmt"
	"sync"
)

func main() {
	// maps()
	// iterateAndWrite()
	slice()
}

func iterateAndWrite() {
	m := map[int]int{}

	go func() {
		for key := range m {
			_ = key
		}
	}()
	for i := 0; i < 10000; i++ {
		m[i] = i
	}

}

func maps() {
	m := map[int]int{}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			m[i] = i
		}(i)
	}
	fmt.Println(len(m))
}

func slice() {
	s := make([]int, 0, 10000)
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			s = append(s, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(s))
}
