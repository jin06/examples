package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond)
	ticker2 := time.NewTicker(time.Millisecond * 2)
	var a, b int
	for {
		select {
		case <-ticker.C:
			{
				a++
				time.Sleep(time.Millisecond * 10)
				fmt.Println("a:", a)
			}
		case <-ticker2.C:
			{
				b++
				time.Sleep(time.Millisecond * 10)
				fmt.Println("b:", b)
			}
		}
	}

}

func main3() {
	// 错误示例代码，time.Tick 每次都生成新的ticker，如果不释放会造成资源浪费
	closed := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				fmt.Println(1)
			case <-closed:
				return
			}
		}
	}()
	time.AfterFunc(time.Second*3, func() { close(closed) })
	select {}
}

func main2() {
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()
	for t := range ticker.C {
		fmt.Println(t)
	}
}

func main1() {
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
	}()
	// time.Tick(time.Second) 仅仅当是程序最后再使用
	<-time.After(time.Second)
}
