package main

import (
	"errors"
	"fmt"
)

//类型简写
func intSum(x, y int) int {
	return x + y
}

//可变参数是指函数数量不固定
func intSum2(x ...int) int {
	fmt.Println(x) //X是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//返回值
//多返回值
func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//返回值补充
//当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

func someFunc(x string) []int {
	if x == "" {
		return nil
	}
	return nil
}

//定义全部变量Num
var num int64 = 10

//如果定义局部变量和全局变量重名，优先访问局部变量
func testGlobalVar() {
	fmt.Println("num\n", num)
}

//语句块定义的变量
func testLocalVar2(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		z := 100 //变量只在if语句、for语句块、switch语句块生效
		fmt.Println(z)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在for语句块中生效
	}
}

//用type关键字来定义一个函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//高阶函数分为函数作为参数和函数作为返回值两部分
//函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return x + y + op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//go语言函数内部不能再想之前那样定义函数、只能定义匿名函数

func main() {
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret4)
	testGlobalVar()
	var c calculation
	c = add
	fmt.Println(c(1, 2))

	f := add
	fmt.Println(f(10, 20))

	ret2 := calc3(10, 20, add)
	fmt.Println(ret2) //30

	//将匿名函数保存到变量
	add2 := func(x, y int) {
		fmt.Println(x, y)
	}
	add2(10, 20) //通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x, y)
	}(10, 30)
	//匿名函数多用于闭包和回调函数

}
