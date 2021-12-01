//Go语言中的值类型（int、float、bool、string、array、struct
//都有对应的指针类型，如：*int、*int64、*string等。
package main

import "fmt"

//ptr:用于接收地址的变量
//v:代表被取地址的变量，类型为T
func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}
func main() {
	a := 100
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
	//否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。分配内存，就引出来今天的new和make

	d := new(int)
	e := new(bool)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Println(*d)
	fmt.Println(*e)

	var g *int
	g = new(int)
	*g = 1000
	fmt.Println(*g)

	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型
	//因为这三种类型就是引用类型，所以就没有必要返回他们的指针了]

	//make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
	var h map[string]int
	h = make(map[string]int, 10)
	h["蒙娜丽莎"] = 100
	fmt.Println(h)
	//new 用于类型的内存分配，并且内存对应的值为类型零值，返回指向类型的指针

}
