package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//go没有类可以为结构体类型定义方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//接口类型时由一组方法签名定义的集合，接口类型的变量可保存实现了这些方法的值
type Aber interface {
	Abs() float64
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//使用指针接受者原因是首先方法能修改其接收者指向的值，其次，可以避免在每次
//调用法师是复制该值，若值的类型为大型结构体时没这样做会更加高效
func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	var a Aber
	f := MyFloat(-math.Sqrt2)
	h := Vertex{3, 4}
	a = f  // Myfloat 实现了Abser
	a = &h //*Vertex实现了Abser
	a = h  //这一行，h是一个Vertex而不是*Vertex所有没有实现Abser
	fmt.Println(a.Abs())
}
