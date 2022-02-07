// Hello World
/*
1. 每个文件都必须归属于一个包
2. func 函数
3. go build 编译为.exe文件
   go run hello.go 运行程序
*/

package main  //打包
import "fmt"  // 引入一个包

func main()  {
	fmt.Println("Hello World")
}