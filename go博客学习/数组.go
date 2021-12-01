package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}
func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {

	var testArray [3]int        //数组会初始化int类型的零值
	var numArray = [3]int{1, 2} //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	//使用指定索引值的方式来初始化数组
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)
	//数组的遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	// for range遍历
	for index, city := range cityArray {
		fmt.Println(index, city)
	}
	//多维数组
	b := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(b)
	fmt.Println(b[2][1])
	//多维数组只有第一层可以使用...来让编译器推导数组长度
	c := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(c)

	//数组是值类型
	d := [3]int{10, 20, 40}
	modifyArray(d)
	fmt.Println(d)
	e := [3][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	modifyArray2(e)
	fmt.Println(e)
	//[n]*T表示指针数组，*[n]T表示数组指针

}
