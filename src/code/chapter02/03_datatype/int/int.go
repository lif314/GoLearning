// 基本数类型使用
package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main(){
	// 整数类型
	// 有符号
	// var i int8 = 129 // constant 129 overflows int8
	var i int8 = 127
	fmt.Println("i=", i)
	// 无符号
	var ui uint8 = 129
	fmt.Println("ui=", ui)

	// int, uint 与系统一致
	// rune = int32, byte = uint8

	// 如何在程序中查看某个变量的字节大小和数据类型
	fmt.Printf("i 的数据类型是 %T\n", i)// 格式化输出
	fmt.Printf("i 的字节大小是 %d", unsafe.Sizeof(i))// 格式化输出
}
