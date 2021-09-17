package main

//隐式接口从接口的实现中解耦定义，这样接口可以是先在任何包
//无需在每一个实现上增加新的接口名称
//隐式接口通过实现一个接口的所有方法来实现该接口
//无需专门显示声明，也就没"implements"
import (
	"fmt"
	"math"
)

//即便接口内的具体值为nil，方法仍然会被nil接收者调用
//在一些语言中这回触发一些空指针异常，但go会优雅的处理
//接口也是值，可以向其他值一样传递，也可作为函数的参数和返回值
type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}
func (f F) M() {
	fmt.Println(f)
}

type F float64

func main() {
	var i I = T{"hello"}
	i.M()
	var k I

	k = &T{"Hello"}
	describe(k)
	k.M()

	k = F(math.Pi)
	describe(k)
	k.M()
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
