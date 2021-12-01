package main

//sync.WaitGroup来实现goroutine的同步
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//并发：同一时间段内执行多个任务
//并行：同一时刻执行多个任务

//Go语言的并发通过goroutine实现。goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。
func hello() {
	fmt.Println("Hello Goroutine!")
}
func hello1(i int) {
	defer wg.Done() //goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}

}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}
func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(10 * time.Second) //所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.sleep
	//是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。

	for i := 0; i < 10; i++ {
		wg.Add(1) //	启动一个goroutine登记 +1
		go hello1(i)
	}
	wg.Wait() //等待所有登记的goroutine都结束
	/*goroutine和OS线程是多对多的关系，即m:n，其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池，
	不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。
	 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。*/

	runtime.GOMAXPROCS(2) //设置当前程序并发时占用的CPU逻辑核心数,1.5版本后默认使用全部CPU逻辑核心数
	go a()                //如果设置只有一个逻辑核心就是做完一个任务再做另一个任务
	go b()
	time.Sleep(time.Second)

}
