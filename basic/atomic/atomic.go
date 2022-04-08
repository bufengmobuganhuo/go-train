package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现一个原子型变量
type atomicInt struct {
	value int
	// 定义一个锁
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		// 加锁
		a.lock.Lock()
		// 只对函数体内部起作用
		defer a.lock.Unlock()
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get())
}
