package main

import "fmt"

// 声明变量的表达式：var 变量名字 类型 = 表达式
//类型和初始表达式可以省略一个。省略类型，会根据初始表达式，自动判断类型。省略初始表达式，会赋予零值，
//数值类型零值为0，布尔类型false，字符串类型 空字符串，接口或引用类型 nil，结构体零值为各个元素上的零值。
func main() {
	var s string
	fmt.Println(s) // ""
}
