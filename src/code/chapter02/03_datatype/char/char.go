package main

import "fmt"

func main(){
	var a byte = 'a' // 使用单引号
	var b byte = '0' // 使用单引号
	// var a byte = "a" // cannot use "a" (type untyped string) as type byte in assignment
	//  输出的是对应的ASCII码值
	fmt.Println("a=", a, "\nb=", b)
	// 输出对应的字符 -- 格式化输出
	fmt.Printf("a=%c \nb=%c", a, b)

	// 汉字溢出
	// var over byte = '李' // constant 26446 overflows byte
	var over int = '李'
	fmt.Printf("\nover=%c", over)

	// 码值运算
	fmt.Printf("\nover=%c", over+100)

	
	// 双引号是string类型
}