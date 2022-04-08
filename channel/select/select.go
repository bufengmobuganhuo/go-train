package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for n := range c {
			time.Sleep(time.Second)
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}()
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	// 定义一个slice，负责接收channel的数据
	var values []int
	// 10秒后会向channel发送一个信号
	tm := time.After(10 * time.Second)
	// 每隔500毫秒会发送一个数据
	tick := time.Tick(500 * time.Millisecond)
	for {
		var activeWorker chan<- int
		var activeValue int
		// >0 说明接收到数据
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		// 谁先收到了就用谁的
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			// values接收到的数据发送给worker
		case activeWorker <- activeValue:
			// 向channel发送成功后，从slice中去除value
			values = values[1:]
			// 在800毫秒内没有数据生成，则会打印（800毫秒内，上面的三个case都没有命中）
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
			// 每隔500毫秒看一下队列长度
		case <-tick:
			fmt.Println("queue length =", len(values))
			// 接收到信号后，说明10秒定时结束
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}
