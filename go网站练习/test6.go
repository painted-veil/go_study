package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//映射将健映射到值
//映射的零值为nil,nil映射没有值，也不能添加健
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) //
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	var f = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(f)
	var u = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(u)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//函数也是值可以像其他值一样传递
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	//函数的闭包，Go函数可以是是闭包，闭包是一个函数值
	//引用了其函数体之外的函数
	//例如函数adder返回一个闭包每个闭包都被绑定在各自的sum变量上
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
