package main

/*
#include <stdio.h>

// C 函数
void hello_from_c() {
    printf("Hello from C!\n");
}
*/
import "C"
import "fmt"

func main() {
	// 调用 C 函数
	C.hello_from_c()

	// Go 代码
	fmt.Println("Hello from Go!")
}
