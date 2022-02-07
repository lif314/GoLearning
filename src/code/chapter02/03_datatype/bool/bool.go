package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var flag bool = true
	fmt.Println("flag=", flag)
	fmt.Println("byte of flag: ", unsafe.Sizeof(flag))
}