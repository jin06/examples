package main

import (
	"fmt"

	"github.com/MauriceGit/skiplist"
)

type Element int

// Implement the interface used in skiplist
func (e Element) ExtractKey() float64 {
	return float64(e)
}
func (e Element) String() string {
	return fmt.Sprintf("%03d", e)
}

type iterator struct {
	skiplist.SkipList
}

func (iter *iterator) f2(yield func(float64, string) bool) {
	last := iter.GetSmallestNode()
	for i := 0; i < iter.GetNodeCount(); i++ {
		if !yield(last.GetValue().ExtractKey(), last.GetValue().String()) {
			return
		}
		last = iter.Next(last)
	}
}

func main() {
	list := iterator{skiplist.New()}
	list.Insert(Element(101))
	list.Insert(Element(3))
	list.Insert(Element(19))
	list.Insert(Element(4))
	list.Insert(Element(89))

	for k, v := range list.f2 {
		fmt.Println(k, v)
	}
}
