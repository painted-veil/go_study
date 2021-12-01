package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x) //可以获得任意值的类型对象
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

//通过发射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64: //// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转化
		fmt.Printf("type is inte64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))

	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}

}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //// 反射中使用 Elem()方法获取指针对应的值
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	reflectValue(a)

	var b int64 = 100
	reflectType(b)
	reflectValue(b)

	var c rune
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "黄河长江",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d)
	reflectType(e)

	f := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", f)

	var g int64 = 100
	reflectSetValue2(g)
	fmt.Println(g)

	var k *int
	fmt.Println("var k *int IsNil:", reflect.ValueOf(k).IsNil())
	//nil值
	fmt.Println("nil isValid", reflect.ValueOf(nil).IsValid())

	//实例化一个匿名结构体
	p := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(p).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(p).MethodByName("abc").IsValid())
	//map
	u := map[string]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(u).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))

	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

//基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
//大量使用反射的代码通常难以理解。
//反射的性能低下，基于反射实现的代码通常比正常运行代码运行速度慢一到两个数量级
