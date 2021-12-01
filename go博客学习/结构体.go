package main

import (
	"fmt"
	"unsafe"
)

//type关键字来定义自定义类型。
type NewInt int

//类型别名规定,只是int一个别名，本质上是同一个类型
type MyInt = int

type person struct {
	name, city string
	age        int64
}

//任何类型都可以拥有方法
func (m NewInt) SayHello() {
	fmt.Println("我是一个Int")
}

//构造函数，自己实现
func NewPerson(name, city string, age int64) *person {
	return &person{
		name: name, city: city, age: age,
	}
}

//方法和接收者
func (p person) Dream() {
	fmt.Printf("%s的梦想和是学好Go语言！\n", p.name)
}

//使用指针接收者

//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
func (p *person) SetAge(newAge int64) {
	p.age = newAge
}

//什么时候需要使用指针类型接收者
//1、需要修改接收者中的值
//2、接受者是拷贝代价比较大的的大对象
//3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者

//rune和byte就是类型别名，他们的定义如下
//type byte = uint8
//type rune = int32

//一个结构体可以嵌套包含另一个结构体或结构体指针
type Address struct {
	Province string
	City     string
}
type User struct {
	Name    string
	Gender  string
	Address //Address
}

func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	var p1 person
	p1.name = "黄河"
	p1.city = "中国"
	p1.age = 5000
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1 =%#v\n", p1)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//创建指针类型结构体

	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2) //p2是一个结构指针
	p2.name = "小王子"
	p2.age = 28
	p2.city = "奥匈帝国"
	fmt.Printf("p2=%#v\n", p2)

	//取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3 = %#v\n", p3)
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3 = %#v\n", p3)

	//没有初始化的结构体，其成员变量都是对应其类型的零值

	//使用键值对初始化
	p5 := person{
		name: "大笨蛋",
		city: "美国",
		age:  12,
	}
	fmt.Printf("p5=%#v\n", p5)
	//也可以对结构体指针进行键值对初始化
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)
	//当某些字段没有初始化值的时候，该字段可以不写，此时没有指定初始值的字段就是该类型的零值

	//初始化结构体的时候可以简写，也就是初始化的时候不写键
	p8 := &person{
		"小兵",
		"广州",
		30,
	}
	fmt.Printf("p8=%#v\n", p8)

	//结构体占用一块连续的内存
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	//空结构体是不占用空间的
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))

	//调用构造函数
	p9 := NewPerson("库里", "旧金山", 33)
	fmt.Printf("%#v\n", p9)
	//方法与函数的区别是函数不属于任何类类型，方法属于特定的类型
	p9.Dream()

	var m1 NewInt
	m1.SayHello()
	//非本地类型不能定义方法，我们不能给别的包的类型定义方法。

	//结构体的匿名字段
	pp1 := person{
		"利拉德",
		"波特兰",
		31,
	}
	fmt.Printf("%#v\n", pp1)
	fmt.Println(pp1.name, pp1.city, pp1.age)

	//嵌套结构体
	u1 := User{
		Name:   "无敌",
		Gender: "女",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("user1=%#v\n", u1)

	//也可以用嵌套匿名字段
	var u2 User
	u2.Name = "小黄鱼"
	u2.Gender = "雄"
	u2.Address.Province = "尼加拉瓜" //匿名字段默认使用类型名作为字段名
	u2.City = "不知道"              //注意这里Address如果在结构体中后面有名字则不能匿名
	fmt.Printf("user2=%#v\n", u2)
}
