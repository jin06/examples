package main

import (
	"fmt"
	"unsafe"
)

var emptyStruct struct{}

var nilVar *int

func testEmptyStruct() {
	fmt.Println("test empty struct")
	fmt.Println("empty struct size: ", unsafe.Sizeof(emptyStruct))
	fmt.Println("nil varible size: ", unsafe.Sizeof(nilVar))

}

type runner interface {
	run()
}

type defaultRunner struct{}

func (defaultRunner) run() {
	fmt.Println("not implement")
}
