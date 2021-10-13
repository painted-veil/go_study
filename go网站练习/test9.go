package main

import "fmt"

//nil接口既不保存值也保存具体类型
//如果为Nil接口调用方法会产生运行时错误，因为接口的元组未包含能够指明该调用哪个具体方法
//指定了零个方法的接口值被称为空接口
type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i) //注意这里printf和printf格式不一样
}
