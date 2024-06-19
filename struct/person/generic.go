package person

import (
	"encoding/json"
	"fmt"
)

type PageList[T any] struct {
	Total   int
	Current int
	Size    int
	List    []T
}

func TestGeneric() {
	type row struct {
		Name  string
		Value string
	}

	type row2 struct {
		Age     int
		Content string
	}

	p1 := PageList[row]{
		List: []row{
			row{Name: "l", Value: "ll"},
			row{Name: "l", Value: "ll"},
			row{Name: "l", Value: "ll"},
		},
	}
	p2 := PageList[row2]{
		List: []row2{
			row2{Age: 2, Content: "hello"},
			row2{Age: 2, Content: "hello"},
			row2{Age: 2, Content: "hello"},
			row2{Age: 2, Content: "hello"},
		},
	}

	b, _ := json.Marshal(p1)
	fmt.Println(string(b))
	b, _ = json.Marshal(p2)
	fmt.Println(string(b))
}
