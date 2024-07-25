/*
	func(func(K,V) bool)
*/

package showcase

import (
	"fmt"
	"testing"
)

func f2(yield func(k int, v string) bool) {
	for i := 0; i < 5; i++ {
		if !yield(i, fmt.Sprintf("I'm %v", i)) {
			return
		}
	}
}

func TestRev2(t *testing.T) {
	for k, v := range f2 {
		fmt.Printf("K: %v, V: %v \n", k, v)
	}
}

func TestBreak(t *testing.T) {
	for k, v := range f2 {
		fmt.Printf("K: %v, V: %v \n", k, v)
		if k == 3 {
			break
		}
	}
}
