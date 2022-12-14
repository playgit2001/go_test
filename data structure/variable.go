package data_structure

import "fmt"

// 声明变量的表达式：var 变量名字 类型 = 表达式
var s1 string = "1234"

//类型和初始表达式可以省略一个。省略类型，会根据初始表达式，自动判断类型。省略初始表达式，会赋予零值，
//数值类型零值为0，布尔类型false，字符串类型 空字符串，接口或引用类型 nil，结构体零值为各个元素上的零值。
func main() {
	var s string
	fmt.Println(s) // ""
	fmt.Println(s1)
	var t1, t2, t3 = 1, 2, 3 //可以声明一组变量
	var t4, t5 int
	fmt.Printf("%d,%d,\n", t1+t2+t3, t4+t5)
	var s2 = Tostring() //变量也能通过一个函数来进行初始化
	fmt.Println(s2)
	s3 := 123 //简短声明。:=是一个变量声明语句，而=是一个赋值操作
	fmt.Println(s3)
	//这里有一个比较微妙的地方：简短变量声明左边的变量可能并不是全部都是刚刚声明的
	//如果有一些已经在相同的词法域声明过了（§2.7），那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。
	//简短变量声明语句中必须至少要声明一个新的变量

	//指针
	//对于每个变量，都有保存它的内存地址，我们用指针对应它的内存地址
	//对于一个变量x，通过它的指针可以读取和更新它的值，可以说这个指针指向了变量x
	var x = 0
	x1 := &x //x1指向x，&符号获取地址
	*x1 = 5  //*符号，提取变量。
	fmt.Println(*x1, x)

	//对于聚合类型，他的每一个元素都是变量，所以可以被取地址
	//变量有时候被称为可寻址的值。即使变量由表达式临时生成，那么表达式也必须能接受&取地址操作
	//任何类型的指针的零值都是nil。如果p指向某个有效变量，那么p != nil测试为真。
	//指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
	var x2, y int
	fmt.Println(&x2 == &x, &x2 == &y, &x2 == nil)

	//在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v.
	//在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
	var p = f()
	fmt.Println(p == f())

	//当函数使用指针作为参数时，若函数体里面通过指针进行了修改，那么当函数体结束时，指针指向的变量保存函数体内被修改后的状态
	var p1 = 10
	p2 := &p1
	tozhen(p2)
	fmt.Println(p1)
	//指针相当于变量的别名，我们可以不通过变量而是别名来访问变量。
	//但是这也是一把双刃剑，如果需要找到该变量的所有访问者，就需要找到所有的别名

	//内建函数new，new(T)将创建一个T类型的匿名变量，并返回该指针
	p3 := new(int) //new(int)相当于 var p int;return &p.
	fmt.Println(*p3)
	//new每次返回一个新地址，但是如果时大小为0的类型，可以两个返回值相同（依赖语言的具体实现）至少对于go语言是如此
	fmt.Println(new(int) == new(int))

	//变量的生命周期
	//变量的生命周期指的是在程序运行期间变量有效存在的时间段。对于在包一级声明的变量来说
	//它们的生命周期和整个程序的运行周期是一致的。而相比之下，局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，
	//然后变量的存储空间可能被回收。函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。
	//具体细节不想bb
}
func tozhen(p *int) {
	*p = 5
}
func Tostring() string {
	return "hello"
}
func f() *int {
	v := 1
	return &v
}
