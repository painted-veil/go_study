package main

import (
	"fmt"
)

func main() {
	var a []string     //声明一个字符串切片
	var b = []string{} //声明一个整形切片并初始化
	var c = []bool{false, true}
	//声明一个布尔切片并初始化
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil) //切片是引用类型，不支持比较，只能和nil比较
	//切片用用自己长度和容量 cap()求容量，len()求长度
	e := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("e:%v len(e):%v cap(e):%v\n", e, len(e), cap(e))
	s := e[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//索引的上限是cap(s)而不是len(s)
	s2 := s[3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	//对于数组，或切片a注意不能是字符串支持完整的切片表达式a[low : high : max]
	t := e[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
	//使用make创造函数切片make([]T, size, cap)
	d := make([]int, 2, 10)
	fmt.Printf("d:%v len(d):%v", d, len(d))

	//切片的本质就是对底层数组的封装底层数组的指针、切片的长度（len）和切片的容量（cap）
	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

	//切片的赋值拷贝
	g1 := make([]int, 3)
	g2 := g1
	g2[0] = 100
	fmt.Println(g1)
	fmt.Println(g2)
	//切片遍历
	w := []int{1, 3, 5}
	for i := 0; i < len(w); i++ {
		fmt.Println(i, w[i])
	}
	for index, value := range w {
		fmt.Println(index, value)
	}

	//切片添加元素
	var f []int
	f = append(f, 1)
	f = append(f, 2, 3, 4)
	f2 := []int{5, 6, 7}
	f = append(f, f2...)
	fmt.Println(f)

	//append函数直接使用不需要初始化,且支持一次性追加多个函数
	s4 := []int{}
	s4 = append(s4, 1, 2, 3)
	fmt.Println(s4)
	var s5 = make([]int, 0, 3) // 这里make要初始化这里 0改为5则表示5个零填充
	s5 = append(s5, 1, 2, 3)
	fmt.Println(s5)
	//每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	//当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	//“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	//使用copy函数复制切片
	//copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	c2 := []int{1, 2, 3, 4, 5}
	c3 := make([]int, 5, 5)
	copy(c3, c2) //如果直接相等则表示同一个地址
	c3[0] = 1000
	fmt.Println(c2)
	fmt.Println(c3)
	//从切片删除元素
	c4 := []int{3, 4, 5, 6, 7, 8, 2, 1}
	//删除索引为2的元素
	c4 = append(c4[:2], c4[3:]...)
	fmt.Println(c4)

}
