package main

import "fmt"

func f2(yield func(int, string) bool) {
	for i := 0; i < 10; i++ {
		if !yield(i, fmt.Sprintf("I'm %d ", i)) {
			fmt.Println("ok break")
			return
		}
	}
}

func main() {
	for k, v := range f2 {
		if k == 7 {
			fmt.Println("go break")
			break
		}
		fmt.Println(v)
	}
}
