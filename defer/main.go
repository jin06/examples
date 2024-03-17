package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	// fn1()
	// returnerr(-1)
	fnpanic()
}
func sort() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}

func file() {
	f, err := os.Open("lang.db")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		l, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(l))
	}
	//do f.Close()
}

func fileWithoutDefer() {
	f, err := os.Open("lang.db")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	for {
		l, _, err := buf.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(l))
	}
	f.Close()
}

var l sync.Mutex
var arr []int

func loopLock() {
	for i := 0; i < 10000; i++ {
		go func() {
			lock()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(len(arr))
}

func lock() {
	l.Lock()
	defer l.Unlock()
	arr = append(arr, 1)
}

func handleErr(err error) {
	fmt.Println("defer")
	if err != nil {
		fmt.Println(err)
	}
}

func returnerr(i int) (err error) {
	// defer function 参数在加入延迟调用栈时已经确定的
	defer func() {
		handleErr(err)
	}()
	if i < 0 {
		err = errors.New("i < 0")
		return
	}

	if i == 0 {
		err = errors.New("i == 0")
		return
	}
	fmt.Println(i)
	return nil
}

func fn3() {
	fmt.Println(3)
}

func fn2() {
	defer fn3()
	fmt.Println(2)
}

func fn1() {
	defer fn2()
	fmt.Println(1)
}

func fnpanic() {
	defer fmt.Println("out")

	fmt.Println(1)

	panic(2)
	defer fmt.Println("after 2")
	fmt.Println("in")
}
