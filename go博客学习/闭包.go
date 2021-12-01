package main

import "fmt"

//闭包 =函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}

}

func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))

	f1 := adder()
	fmt.Println(f1(40))
	fmt.Println(f1(50))

	f3, f4 := cal(10)
	fmt.Println(f3(1), f4(9))
	fmt.Println(f3(3), f4(4)) //12 8
	fmt.Println(f3(5), f4(6)) //13 7

	//defer语句将其后面跟随的语句进行延迟处理
	//在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
	//	先被defer的语句最后被执行，最后被defer的语句，最先被执行。

	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	//go语句的return在底层并不是原子操作
	//	它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
}
