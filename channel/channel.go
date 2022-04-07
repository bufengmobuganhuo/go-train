package main

import (
	"fmt"
	"time"
)

// chan<- int：表示函数调用者只能向channel发送数据
// <-chan int：相反
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			// 将channel的数据赋值给n
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		// 将channel的数据赋值给n
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func chanDemo() {
	// 定义10个可以发送int的channel
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		// 定义一个goroutine，负责从channel收数据
		go worker(i, channels[i])
	}
	// 向每个channel发送数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	// 向每个channel发送数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func createWorkerDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// 存储返回的channel
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		// 给channel发送数据，注意在发送之前必须定义好接收的goroutine，否则会陷入死锁(发送是阻塞的，发送了数据后会等待别人收数据)
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func bufferredChannel() {
	// 定义缓冲区为3的channel
	c := make(chan int, 3)
	// 因为有缓冲区的存在，可以先不定义goroutine，但是只能发送3个
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	// 关闭channel，只能发送方close
	close(c)
}

func main() {
	//chanDemo()
	//createWorkerDemo()
	//bufferredChannel()
	channelClose()
}
