package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	printSlice(s)
	var g []int
	fmt.Println(g, len(g), cap(g))
	if g == nil {
		fmt.Println("nil!")
	}
	//切片可以用内建函数make来创建，这也是你创建动态数组的方式
	a := make([]int, 5)
	printSlice2("a", a)
	//切片可包含任何类型甚至包括其他的切片
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	var k []int
	printSlice(k)
	k = append(k, 1) //这个切片会按需增长
	printSlice(k)
	k = append(k, 2, 3, 5)
	printSlice(k)
	//for循环的range形式可遍历切片或映射，每次迭代会返回两个值，一个为当前元素的下标，第二个为该下标所对应元素一份副本
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//也可以将表表或值赋予_来忽略他
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
