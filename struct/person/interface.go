package person

import "fmt"

type testinf interface {
	test()
}

type testins1 struct{}

func (t testins1) test() {
	fmt.Println("test1")
}

type testins2 struct{}

func (t testins2) test() {
	fmt.Println("test2")
}

type MyTest struct {
	Instance testinf
}

func (m *MyTest) runTest() {
	m.Instance.test()
}

func TestInterface() {
	m := MyTest{}
	m.Instance = testins1{}

	m.runTest()

	m.Instance = testins2{}

	m.runTest()
}
