package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	for i := 0; i < 100; i++ {
		lock.Lock() //加锁
		x++
		lock.Unlock() //解锁
	}
	wg.Done()
}
func hello() {
	defer wg.Done()
	fmt.Println("hello,goroutine")
}

//读写锁在Go语言中使用sync包中的RWMutex类型。
/*锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。*/

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done!")
	wg.Wait()
	//Go语言中可以使用sync.WaitGroup来实现并发任务的同步。
	//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成
}
