package main

import (
	"fmt"
	"time"
)

//类型选择时一种按孙旭从几个类型断言中选择分支的结构
//一般与switch语句相似，针对给定接口值所存储的值类型进行比较
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

//Stringrt是最普遍的接口之一
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//Go 程序使用 error 值来表示错误状态
//error是一个内建的接口，通常函数会返回一个error值，调用它的的代码应当判断这个错误是否等于nil进行错误处理

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What) //这里是格式化输出
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func main() {
	do(21)
	do("hello")
	do(true)
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
