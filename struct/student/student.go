package student

import "fmt"

// student
type Student struct {
	No      int
	Name    string
	Class   int
	private string
}

func TestPrivate() {
	s := Student{
		No:      1,
		Name:    "tom",
		private: "I have a secret",
	}
	fmt.Println("包内部 testPrivate", s.private)
}

func NewCar() *car {
	return &car{}
}

type car struct {
	Name   string
	Field1 string
	Field2 string
}

func (c *car) Run() {
	fmt.Println(c.Name + " running")
}

func (c car) Run2() {
	fmt.Println(c.Name + " running")
}

func (c *car) Change(newName string) {
	c.Name = newName
}
func (c car) Change2(newName string) {
	c.Name = newName
}

func (car) Table() string {
	return "car"
}

func TestCarTable() {
	c := car{}
	fmt.Println(c.Table())
}

func TestChange() {
	c1 := car{
		Name: "xiaomi su7",
	}
	c2 := car{
		Name: "xiaomi su7",
	}

	c1.Change("tesla model3")
	c2.Change2("tesla model3")
	fmt.Println(c1.Name)
	fmt.Println(c2.Name)
}

func testCar() {
	c := car{
		Field1: "f1",
		Field2: "f2",
	}
	fmt.Println("包内部 testCar", c)
}
