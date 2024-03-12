package main

import "fmt"

func main() {
	// map 小于等于8 ，有相对顺序
	// map 大于8 ，没有
	// mapping := map[any]int{
	// 	"car":  1,
	// 	"rust": 1,
	// 	2:      1,
	// 	"io":   1,
	// 	"echo": 1,
	// 	"pwd":  1,
	// 	6:      1,
	// 	7:      1,
	// 	8:      1,
	// }

	// for i := 0; i < 5; i++ {
	// 	for key := range mapping {
	// 		fmt.Printf("%v ", key)
	// 	}
	// 	fmt.Println()
	// }
	printfmap := func(m map[int]int) {
		for key := range m {
			fmt.Printf("%v ", key)
		}
		fmt.Println()
	}

	m2 := map[int]int{}
	for i := 0; i < 9; i++ {
		m2[i] = i
		printfmap(m2)
	}
	printfmap(m2)
	printfmap(m2)
	printfmap(m2)
	printfmap(m2)
}
