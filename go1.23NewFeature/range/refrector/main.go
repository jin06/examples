package main

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

	res1 := make([]int, 0)
	q.Iterator(func(cid int) {
		res1 = append(res1, cid+1)
	})

	for _, cid := range q.f2 {
		res1 = append(res1, cid+1)
	}

}
