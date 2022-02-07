// 转义字符
package main
import "fmt"

func main(){  // 一个包下只能有一个main()函数
	fmt.Println("Hello \t World") // 制表符
	fmt.Println("Hello \\ World") // \
	fmt.Println("Hello \n World") // 换行
	fmt.Println("Hello \r World") // 回车
}