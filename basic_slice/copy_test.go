package basic_slice

import (
	"reflect"
	"testing"
)

// 切片是不是引用类型
func TestSliceType(t *testing.T) {
	values := make([]int, 2, 2)

	handle := func(data []int) {
		data[0], data[1] = 0, 1
		for i := range 3 {
			data = append(data, i)
		}
	}

	handle(values)
	// 0,0,0,1,2
	typ := reflect.TypeOf(values)
	t.Log(typ)
}

func TestSliceType2(t *testing.T) {
	values := []int{1, 2, 3, 4}

	v2 := values[:2] // 1,2 cap=4

	v3 := values[:2:2] // 1,2 cap=2

	v3 = append(v3, 5)

	t.Log(values, v2, v3)
}

func TestSliceType3(t *testing.T) {
	values := []int{1, 2, 3, 4}

	values2 := values
	values3 := values[1:2]
	t.Logf("%p", values)
	t.Logf("%p", values2)
	t.Logf("%p", values3)
}

// go test -v --run TestCopy
// copy
func TestCopy(t *testing.T) {
	source := []int{1, 2, 3}
	dest := make([]int, 1, 2)
	result := copy(dest, source)
	t.Log("result: ", result)
	t.Log("dest: ", dest)
}

func TestForever(t *testing.T) {
	values := []int{1, 2}

	for _, i := range values {
		t.Log(i)
		values = append(values, i)
	}
	t.Log(values)
}

func TestNil(t *testing.T) {
	var v1 []int
	v2 := []int{}
	v3 := make([]int, 0)

	t.Log(v1 == nil, v2 == nil, v3 == nil)
}
