package main

//go程是由go运行管理轻量级线程
import (
	"fmt"
	"time"
)

//信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
//“箭头”就是数据流的方向。
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步
	//启动一个go程
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	//和映射与切片一样，信道在使用前必须创建

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
	//信道也是可以带缓冲的
	ch := make(chan int, 2) //若缓冲区填满会发生堵塞
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//要发送者可通过 close 关闭一个信道来表示没有需发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
	//只有发送者才能关闭信道接收者不能

	//信道与文件不同，通常情况下无需关闭它们，只有在必须告诉接收者
	//不在有需要发送的值才有必要要关闭
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}
}
