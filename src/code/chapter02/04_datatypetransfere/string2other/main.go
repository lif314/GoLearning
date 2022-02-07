package main

import (
	"fmt"
	"strconv"
)


// String转基本数据类型
// func ParseBool(str string) (value bool, err error)
// func ParseInt(s string, base int, bitSize int) (i int64, err error)
// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
// func ParseFloat(s string, bitSize int) (f float64, err error)
// func Atoi(s string) (i int, err error)

func main(){
	// func ParseBool(str string) (value bool, err error)
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v\n", b, b)

	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	var str1 string = "1232"
	var a int64
	a, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("a type %T a = %v\n", a, a)

	// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	var fstr string = "1212.121"
	var f float64
	f, _ = strconv.ParseFloat(fstr, 64)
	fmt.Printf("f type %T f = %v\n", f, f)

	// func Atoi(s string) (i int, err error)

	// 确保String类型能够转成有效的数据，如不能是”hello“，将会被转为默认数据，如转为Int则结果为0
	var hello string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(hello, 10, 64)
	fmt.Printf("n3 type %T n3 = %v\n", n3, n3)
}