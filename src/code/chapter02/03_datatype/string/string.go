package main

import (
	"fmt"
)

func main(){
	// 基本使用
	var address string = "上海嘉定"
	fmt.Println(address)

	// Go的字符串是不可变的：一旦赋值则不能更改内部元素
	// var str = "hello"
	// str[0] = "q" // cannot assign to str[0] (strings are immutable)

	// 反引号表示法 -- 整体输出
	var str1 string = `abc\nasada
	Hello Wotld
	\t asafafa`
	fmt.Println(str1)

	// 字符串拼接 +
	fmt.Println(str1 + address)

	// 多行+: 需要将+保留在上一行
	var str2 = "hello" + "World" +
	 "Hello"
	fmt.Println(str2)
}
