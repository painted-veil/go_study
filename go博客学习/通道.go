package main

import "fmt"

//Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
//如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
/*
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
*/

//只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

/*对一个关闭的通道再发送值就会导致panic。
对一个关闭的通道进行接收会一直获取值直到通道为空。
对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
关闭一个已经关闭的通道会导致panic。*/

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}
func squarter(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) { //注意格式
	for i := range in {
		fmt.Println(i)
	}
}

//chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
//<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作
func main() {

	var ch chan int //创建通道，通道类型的空值是nil
	fmt.Println(ch)

	ch1 := make(chan int)  //声明通道后需要使用make函数初始化才能使用
	ch2 := make(chan bool) //
	ch3 := make(chan []int)
	fmt.Println(ch1, ch2, ch3)

	//channel操作，发送和接受都使用<-符号

	/*
		ch1 <- 10  //把10发送到ch1中
		x := <-ch1 //从ch1中接收值并赋值给变量x
		<-ch1      //从通道接受值并忽略结果
		fmt.Println("发送成功 ", x)
		close(ch1) //关闭通道
	*/
	ch4 := make(chan int)
	go recv(ch)
	ch4 <- 10
	fmt.Println("发送成功")
	//deadlock错误，因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值
	//面的代码会阻塞在ch1 <- 10这一行代码形成死锁，
	//解决办法是启用一个gorutine去接收值,

	//无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	//使用无缓冲通道进行通信导致发送和接收的goroutine同步化，因此无缓冲通道也被称为同步通道

	//有缓冲通道
	ch5 := make(chan int, 1) //只要通道大0，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量
	ch5 <- 10
	fmt.Println("发送成功")
	close(ch5)

	//for range从通道循环取值
	//两种方式在接收值的时候判断该通道是否被关闭，不过我们通常使用的是for range的方式。使用for range遍历通道，当通道被关闭的时候就会退出for range

	ch11 := make(chan int)
	ch12 := make(chan int)
	//开启goroutine将0-100的数发送到ch11
	go func() {
		for i := 0; i < 10; i++ {
			ch11 <- i

		}
		close(ch11)
	}()
	//开启goroutine从ch11接收值，并将该值的平方发送到ch2
	go func() {
		for {
			i, ok := <-ch11
			if !ok {
				break
			}
			ch12 <- i * i
		}
		close(ch12)
	}()
	//在主goroutine中从ch12中接受值并打印
	for i := range ch12 { //通道关闭后会退出for range循环
		fmt.Println(i)
	}
	//单向通道
	ch6 := make(chan int)
	ch7 := make(chan int)
	go counter(ch6)
	go squarter(ch7, ch6)
	printer(ch6)
	//关闭已经关闭的channel也会引发panic
}
