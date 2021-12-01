package main

import (
	"fmt"
	"math"
)

//遍历字符串
func traversalString(s string) string {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])

	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
	return s
}

//修改字符串,需要先将其转换成[]rune或[]byte，完成后再转换为string
func changeString() {
	s1 := "big"
	byteS1 := []byte(s1) //强制类型转换
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

//类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	//math.Sqrt()接受的参数是float64类型现需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func main() {
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) //表示二进制

	var b int = 077        //八进制以0开头
	fmt.Printf("%o \n", b) //表示八进制

	//十六进制以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	fmt.Printf("%f \n", math.Pi)
	fmt.Printf("%.2f \n", math.Pi) //支持float32和float64两种浮点型
	//复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	//Go 语言中不允许将整型强制转换为布尔型.
	//布尔类型变量的默认值为false

	//字符串的类型为双引号
	s1 := "hello"
	s2 := "world"
	fmt.Println(s1, s2)
	//多行字符串
	s3 := `第一行
		第二行
		第三行
		`
	fmt.Println(s3)
	//uint8 叫Byte类型，代表一个ASCII字符
	//rune类型代表一个utf-8字符
	_ = traversalString(s1)

}
