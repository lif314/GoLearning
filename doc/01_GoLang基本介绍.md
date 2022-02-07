# 第1章 了解GoLang

[TOC]

## 1. Go基本介绍

- **GoLang优势**

区块链、后台服务器、高并发计算

![image-20220206021604416](https://gitee.com/lilinfei314/PicGo/raw/master//pic/202202060216480.png)

- **如何成为合格的程序员**---学习方法

![image-20220206020819887](https://gitee.com/lilinfei314/PicGo/raw/master//pic/202202060208994.png)

- Go**特点**  Go=C + Python

![image-20220206021906085](https://gitee.com/lilinfei314/PicGo/raw/master//pic/202202060219149.png)



## 2. 执行流程

```go
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
```

![image-20220206132208009](https://gitee.com/lilinfei314/PicGo/raw/master//pic/202202061322150.png)

## 3. 执行命令

- 编译

```go
go build
```

- 执行

```go
go run hello.go
```

## 4. 开发注意事项

- 扩展名 `.go`
- Go应用程序的执行入口是main函数
- 严格区分大小写
- Go不需要加`;`，由编译器自动加上
- Go编译器是**一行一行的进行编译**，不能把多条语句写在一行中
- Go语言中声明的变量或者import的包如果没有用到，则会编译出错
- 大括号成对出现
