package examples

import (
	"fmt"

	"github.com/example/struct/student"
)

func TestPrivate() {

	s := student.Student{
		No:   2,
		Name: "Jack",
	}

	fmt.Println("包外部 testPrivate", s)
}

func TestCar() {
}
