//接口是一种类型
package main

import "fmt"

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

//接口嵌套
type animal interface {
	Sayer
	Mover
}
type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

type Dog struct {
	name string
}
type car struct {
	brand string
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func (d Dog) Move() {
	fmt.Println("黄狗拉里奥")
}

func (c car) Move() {
	fmt.Printf("%s速度70迈", c.brand)
}

//空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var x Sayer //声明一个Say类型的变量
	var y Mover
	a := Cat{}
	b := Dog{name: "旺财"}
	x = a //可以把CAT实例化直接赋值给x
	x.Say()
	x = b //可以把dog实例直接赋值给x
	x.Say()
	y = b
	y.Move()

	//值接收者实现接口
	var wangcai = Dog{name: "wangcai"}
	y = wangcai
	var fugui = &Dog{} //富贵是*dog类型
	y = fugui
	y.Move()

	var ca = car{brand: "保时捷"}
	y = ca
	y.Move()
	y = b
	y.Move()

	//实现animal接口
	var an animal
	an = Dog{name: "animal"}
	an.Move()
	an.Say()

	//空接口
	var i interface{}
	s := "hello 长江"
	i = s
	fmt.Printf("type: %T value:%v\n", i, i)

	g := 100
	i = g
	fmt.Printf("type: %T value:%v\n", i, i)

	bo := true
	i = bo
	fmt.Printf("type: %T value:%v\n", i, i)

	//使用空接口可以保存任意值的字典
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

}
