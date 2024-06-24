package main

import (
	"fmt"
)

type Equal struct {
	A string
	B int
}

type NonEqual struct {
	A string
	B map[string]string
}

type ComplexEqual struct {
	A string
	B Equal
}

func testEqual() {
	fmt.Println("test equal")
	ea := Equal{}
	eb := Equal{}
	if ea == eb {
		fmt.Println("ea == eb")
	}
	mapping := map[Equal]bool{}
	mapping[ea] = true

	fmt.Println(mapping)

}
