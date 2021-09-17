package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) { //go的返回值可以被命名，视作定义在函数顶部的变量，没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var e, f int = 1, 2

func main() {
	fmt.Println("My favorite number is", math.Sqrt(16))
	a, b := swap("world", "hello")
	fmt.Println(add(10, 20))
	fmt.Println(a, b)
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(e, f)
	k := 3 //代替var
	fmt.Println(k)

}
