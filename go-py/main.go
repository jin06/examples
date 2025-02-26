package main

/*
#cgo LDFLAGS: -L/usr/local/lib -lpython3.13
#cgo CFLAGS: -I/usr/local/opt/python3.13
#include <Python.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// 初始化 Python 解释器
	C.Py_Initialize()

	// 执行 Python 代码
	code := C.CString(`
def add(a, b):
    return a + b

result = add(5, 7)
print("Result from Python:", result)
`)
	defer C.free(unsafe.Pointer(code))

	C.PyRun_SimpleString(code)

	// 关闭 Python 解释器
	C.Py_Finalize()

	fmt.Println("Python code executed successfully")
}
