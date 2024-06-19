package main

import (
	"fmt"

	"github.com/example/struct/examples"
	"github.com/example/struct/person"
	"github.com/example/struct/student"
)

func main() {
	student.TestPrivate()

	examples.TestPrivate()

	student.TestChange()
	student.TestCarTable()

	c := student.NewCar()
	c.Name = "bmw"

	fmt.Println(c)

	f := person.Father{}
	s := person.Son{
		Father:      f,
		GrandFather: person.GrandFather{},
	}
	fmt.Println(s.GrandFather.Hello())
	fmt.Println(s.Echo("hello"))

	person.TestInterface()
	person.TestGeneric()
}
