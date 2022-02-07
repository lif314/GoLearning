package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	var num1 int = 90
	var num2 float64 = 32.1213
	var b bool = true
	var char byte = 'a'
	var str string // 空str

	// Sprintf根据format参数生成格式化的字符串并返回该字符串。
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %v\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %v\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %v\n", str, str)
	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type %T str = %v\n", str, str)

	// strconv包中函数
	var str1 string
	str1 =  strconv.FormatBool(b)
	fmt.Printf("str type %T str = %v\n", str1, str1)
	str1 = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str = %v\n", str1, str1)
	str1 = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str = %v\n", str1, str1)
}