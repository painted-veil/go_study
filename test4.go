package main

import "fmt"

//go也有结构体
type Vertex struct {
	X int
	Y int
}

var (
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

//go拥有指针。指针保存了值的内存地址
func main() {
	fmt.Println("test4")
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(Vertex{1, 2}) //注意这里是大括号
	v := Vertex{1, 2}
	v.X = 3
	v.Y = 4
	fmt.Println(v.X, v.Y)
	//结构体字段可以通过结构体指针来访问
	o := &v
	o.X = 1e9 //注意是1不是l
	fmt.Println(o.X)
	fmt.Println(v, p, v2, v3)
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//go的切片
	var s []int = primes[1:4]
	fmt.Println(s) //切换不存储任何数据只是描述底层数组中的一段
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	temp1 := names[1:4]
	temp2 := names[1:3]
	fmt.Println(temp1, temp2)
	temp2[0] = "xxx"
	fmt.Println(temp1, temp2)
	fmt.Println(names) //注意names也会因为更改切片元素修改底层数组对应元素
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	k := []struct {
		g int
		f bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true}, //注意这里的逗号如果大括号在下一行则要逗号
	}
	fmt.Println(k)
}
