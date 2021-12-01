package main

import (
	"fmt"
)

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

//结构体的继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s动起来\n", a.name)
}

type Dog struct {
	Feet    int64
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会叫\n", d.name)
}

//结构体方法补充因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
	//如果是p.dreams = dreams，则当修改数组变量时，结构体的元素也会随之改变，只用这样才不会修改原先结构体
}

//结构体标签
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTim

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意这里是嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()

	//结构体中字段大写开头表示可公开访问，小写表示私有(仅在定义当前结构体包中可访问)

	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	data[1] = "不睡觉"
	fmt.Println(p1.dreams)

	/*s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)  不知道Marsshal用不了*/
}
