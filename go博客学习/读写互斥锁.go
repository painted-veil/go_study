package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type singleton struct{}

var instance *singleton
var once sync.Once

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock() //加写锁
	x++
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock() //解写锁
	//lock.Unlock() //解互斥锁
	wg.Done()
}
func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

//并发安全的单例模式
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

//其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

//Go语言中内置的map不是并发安全的。
var m = sync.Map{} // make(map[string]int)这样map并发是不安全的
/*
用make map高并发不安全
func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

	func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
	}

*/
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write() //
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n) //不用set
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
