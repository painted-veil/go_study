package main

import (
	"fmt"
)

func foo() (int, string) {
	return 10, "Q1mi"
}

// 全局变量m
var m = 100

func main() {
	n := 10 //短变量声明:=
	fmt.Println(m, n)
	//--------------------
	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y) //函数外的每个语句都必须以关键字开始
	//_多用于占位符，表示忽略值
	//--------------------
	//常量
	const (
		pi = 3.1415
		e  = 2.7182
	)
	//如果const同时声明了多个常量如果忽略了值表示和上面个的一行的值相同
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)
	//iota是go语言的常量计数器，只能在常量的表达式中使用
	const (
		n4 = iota //iota在const关键字出现时将充值为0，const中每增一行常量使iota计数一次
		n5
		_
		n6
	)
	fmt.Println(n4, n5, n6)
}
