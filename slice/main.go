package main

import "fmt"

/*
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
*/

func main() {
	var s []int

	// 定义切片s，元素为0，内存分配10
	s = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		s = append(s, i+1)
	}
	s2 := s[:5]
	s3 := s[5:]
	fmt.Printf("s len: %d cap: %d data: %v \n", len(s), cap(s), s)
	fmt.Printf("s2 len: %d cap: %d data: %v \n", len(s2), cap(s2), s2)
	fmt.Printf("s3 len: %d cap: %d data: %v \n", len(s3), cap(s3), s3)
	fmt.Printf("%p %p %p %p \n", s, s2, s3, &s[5])

	s2 = append(s2, 1001)
	fmt.Printf("s len: %d cap: %d data: %v \n", len(s), cap(s), s)
	fmt.Printf("s2 len: %d cap: %d data: %v \n", len(s2), cap(s2), s2)
	fmt.Printf("s3 len: %d cap: %d data: %v \n", len(s3), cap(s3), s3)

	// 切片扩容的时候，其他指向该切片底层数组的切片不会发生变化，其他切片还是指向原来的数组
	s3 = append(s3, 2002)
	fmt.Printf("s len: %d cap: %d data: %v \n", len(s), cap(s), s)
	fmt.Printf("s2 len: %d cap: %d data: %v \n", len(s2), cap(s2), s2)
	fmt.Printf("s3 len: %d cap: %d data: %v \n", len(s3), cap(s3), s3)
	fmt.Printf("%p %p %p %p \n", s, s2, s3, &s[5])

	// 	5 -> 10  -> 20
	// 到了阀值后，线性增长
}
