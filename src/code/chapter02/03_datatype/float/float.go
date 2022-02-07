package main

import "fmt"

func main(){
	// 单精度
	var f float32 = 32.1212121212121212
	fmt.Println("f=", f)

	// 双精度
	var ff float64 = 32.1212121212121212
	fmt.Println("ff=", ff)

	// 默认类型
	var num = 1.1
	fmt.Printf("num 的数据类型是 %T\n", num)// 格式化输出
}