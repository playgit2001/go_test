package main

import "fmt"

//在函数外的变量属于包一级，在同个包下的每个源文件中访问
//var variable1 string //变量声明

/*type variable2 struct { //结构声明
	ID   int64
	name string
}*/

const name = 123 //常量声明
func main() {
	//函数内部定义的变量只能在函数内使用
	const freezingF, boilingF = 32.0, 212.0
	fmt.Println("123")
	fmt.Printf("12321")
	fmt.Println("123")
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
	fmt.Printf("1232141241")
	fmt.Printf("123214123121241")
}

//方法包括了一个函数的声明由一个函数名字、参数列表（由函数的调用者提供参数变量的具体值）
//一个可选的返回值列表和包含函数定义的函数体组成。如果函数没有返回值，那么返回值列表是省略的。
//函数从第一条语句开始执行，直到遇到return语句，或者是到了函数的末尾，之后返回到函数的调用者。
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
