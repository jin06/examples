package main

import "fmt"

type Q struct {
	cids []int
}

func (q *Q) Iterator(f func(cid int)) {
	for _, id := range q.cids {
		f(id)
	}
}

func (q *Q) f2(yield func(int, int) bool) {
	for k, v := range q.cids {
		if !yield(k, v) {
			return
		}
	}
}
func main() {
	q := Q{cids: make([]int, 0)}
	for i := 0; i < 10; i++ {
		q.cids = append(q.cids, i)
	}

	q.Iterator(func(cid int) {
		fmt.Printf("send message: 'hello' to cid: %v", cid)
	})

	for _, cid := range q.f2 {
		fmt.Printf("send message: 'hello' to cid: %v", cid)
	}

}
