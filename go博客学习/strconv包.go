package main

import (
	"fmt"
	"strconv"
)

func main() {

	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}

	//iota 函数用于将Int类型转换为对应字符串表示
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	//Parse系列函数，用于转换字符串为给定类型的值
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64) //类似parseint 但不接受正负号，用于无符号整型
	fmt.Println(b, f, i, u)
	//format系列函数，Format系列函数实现了将给定类型数据格式化为string类型数据的功能
	s4 := strconv.FormatBool(true)
	s5 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s6 := strconv.FormatInt(-2, 16)
	s7 := strconv.FormatUint(2, 16)
	fmt.Println(s4, s5, s6, s7)
	//isprint返回一个字符是否可以打印，canbackquote返回字符串是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。
}
