package main

import "fmt"

func main() {
	var inte interface{}
	describe(inte)
	inte = 42
	describe(inte)
	inte = "hello"
	describe(inte)
	//类型断言,提供访问接口值底层具体值的方式
	var g interface{} = "hello"
	//该语句断言接口值g保存了具体类型string，并将其底层类型string 的值赋予变量
	//若s未保存g类型的值,该语句会触发一个恐慌
	//如果为了判断接口值是否保存了一个特定的类型，类型断言可以返回两个值：底层值以及一个报告断言是否成功的布尔值
	s := g.(string)
	fmt.Println(s)
	s, ok := g.(string)
	fmt.Println(s, ok)

	f, ok := g.(float64)
	fmt.Println(f, ok)
}
func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
