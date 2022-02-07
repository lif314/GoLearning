package main

import (
	"fmt"
)

func main(){
	var i float32 = 10.23
	var int_8 = int8(i)
	var double float64 = float64(i)
	fmt.Println("int_8=", int_8, "\ndouble=", double)


	// 编译问题
	var n1 int32 = 32
	var n2 int64
	var n3 int8

	// cannot use n1 + 20 (type int32) as type int64 in assignment
	// n2 = n1 + 20
	// cannot use n1 + 20 (type int32) as type int8 in assignment
	// n3 = n1 + 20
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println(n2, n3)
}