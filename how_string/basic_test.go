package howstring

import (
	"fmt"
	"os"
	"testing"
)

var a string = "1"
var b string = "a"
var c = "😂"
var d = "张"
var e = "1a😂张"

func TestLen(t *testing.T) {
	fmt.Fprintf(os.Stdout, "len(%s)=%d\n", a, len(a))
	fmt.Fprintf(os.Stdout, "len(%s)=%d\n", b, len(b))
	fmt.Fprintf(os.Stdout, "len(%s)=%d\n", c, len(c))
	fmt.Fprintf(os.Stdout, "len(%s)=%d\n", d, len(d))
	fmt.Fprintf(os.Stdout, "len(%s)=%d\n", e, len(e))
}

func TestRawBytes(t *testing.T) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("十进制：%d\n", a[i])
		fmt.Printf("二进制: %b\n", a[i])
	}
	for i := 0; i < len(d); i++ {
		fmt.Printf("十进制：%d\n", d[i])
		fmt.Printf("二进制：%b\n", d[i])
	}
}

func TestGoSourceCode(t *testing.T) {
	// 😂  := c
	中文 := d

	fmt.Println(中文)
	中文方法(1)
}

func 中文方法(a int) {
	fmt.Printf("中文方法（%d）\n", a)
}

func TestRune(t *testing.T) {
	fmt.Println(d)
	fmt.Printf("%x\n", d)
	for _, v := range d {
		fmt.Printf("%x\n", v)
	}
}
