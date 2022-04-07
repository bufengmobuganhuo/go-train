package main

import (
	"fmt"
	"sync"
)

func doWork(id int, worker worker) {
	for n := range worker.in {
		// 将channel的数据赋值给n
		fmt.Printf("Worker %d received %c\n", id, n)
		// 一个任务执行完成
		worker.done()
	}
}

type worker struct {
	in chan int
	// 表示任务结束的函数
	done func()
}

// chan<- int：表示函数调用者只能向channel发送数据
// <-chan int：相反
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	// 有20个任务需要等待完成
	wg.Add(20)
	// 定义10个可以发送int的channel
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	// 向每个channel发送数据
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// 向每个channel发送数据
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	// 等待执行完成
	wg.Wait()
}

func main() {
	chanDemo()
}
