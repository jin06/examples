package main

import (
	"errors"
	"fmt"
)

func TestClosure() {
	var err error

	defer func() {
		if err != nil {
			// handle
			fmt.Println(err)
		}
	}()

	err = fn()
}

func fn() error {
	return errors.New("fn error")
}

func main() {
	TestClosure()
}
