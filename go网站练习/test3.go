package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //if语句可以在条件标识前执行一个简单语句
		return v
	} else { //这里的else要跟在if最后一个大括号后面
		fmt.Printf("%g >=%g\n", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum = sum + i
	}
	fmt.Println(sum)
	sum2 := 1
	for sum2 <= 100 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	//go中的while用for代替
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)
	//for{
	//这是go的无限循环
	//}
	fmt.Println(pow(3, 3, 10))
	fmt.Println("go runs on")       //go自动提供case后面所需的case
	switch os := runtime.GOOS; os { //switch的case无需为常量，且取值不必为整数
	case "drawin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. \n", os)
	}
	fmt.Println("when's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")
	}
	//没有条件的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	//defer语句将函数推迟到外层函数返回之后执行，推迟调用的函数其参数会立即求值
	//但知道外层函数返回前该函数都不会被调用
	defer fmt.Println("world")
	fmt.Println("hello") //推迟函数的调用会被压入一个栈中，当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}
