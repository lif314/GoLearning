package main

import (
	"fmt"
)

func main(){
	// 基本数据类型的内存布局
	var num int = 10
	fmt.Printf("num的地址：%p", &num)
	// num的地址：0xc0000aa058

	var ptr *int = &num
	fmt.Printf("\nptr的值：%p", ptr)
	fmt.Printf("\nptr的地址：%p", &ptr)
	fmt.Printf("\nptr指向的值：%d", *ptr)
}