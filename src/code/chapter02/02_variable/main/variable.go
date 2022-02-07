// 变量

package main
import "fmt"

// 全局变量
var gn1 = 100
var gn2 = 200

// 一次性全局声明
var (
	gnn1 = 100
	gnn2 = 300
)

func main() {
	// - 指定变量类型，声明后若不赋值，则使用默认值
	var i int
	// 变量不能只定义而不调用
	fmt.Println("i= ", i)

	// - 类型推导：根据值自行判定变量类型
	var num = 10.01
	fmt.Println("num =", num)

	// - 省略`var`，注意`:=`左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 此时应该使用 :=
	name  := "var"
	// var name string
	// name = "var"
	fmt.Println(name)

	// 多变量声明 -- 同类型
	var n1, n2, n3 int 
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 不同类型
	var a, str, c = 100, "tom", 888
	// a, str, c := 100, "tom", 888
	fmt.Println("a=", a, "name=", str, "c=", c)
}