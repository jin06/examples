package main

import "fmt"

func main() {
	fmt.Println("func Add")
	fmt.Println(Add(1, 2))
	fmt.Println(Add[int](2, 2))
	fmt.Println(Add(1.1, 2.2))

	fmt.Println("func Add2")
	fmt.Println(Add2(1, 2))
	fmt.Printf("%+v \n", Add2(float64(1.0), float64(2.5)))

	fmt.Println("func Double")
	fmt.Println(Double([]int{1, 2, 3, 4}))

	fmt.Println("func Index")
	fmt.Println(Index([]int{3, 41, 5, 1}, 5))

	fmt.Println("struct")

	q := Queue[int]{Data: []int{}}
	q.Push(1, 2, 3, 4)
	q.Pop(2)
	fmt.Println(q.Data)
}
