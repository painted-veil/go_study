package main

//image包定义了Image接口
import (
	"fmt"
	"image"
)

//Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
