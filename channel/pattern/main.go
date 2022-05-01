package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	// 生成器，发送消息
	go func() {
		i := 0
		for {
			// sleep 2000毫秒左右
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s generate message: %d", name, i)
			i++
		}
	}()
	return c
}

// 空struct比bool更省空间
func msgGen2(name string, done chan struct{}) chan string {
	c := make(chan string)
	// 生成器，发送消息
	go func() {
		i := 0
		for {
			select {
			// 2秒左右后会发送一个信号，相当于每个每隔2秒左右执行一次
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s generate message: %d", name, i)
				// 收到结束的信号，结束循环
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		// ch变量只有一份，所以后面的channel会覆盖前面的，所以这里需要拷贝一份传过去
		chCopy := ch
		go func() {
			// 将c1的数据取出送给c
			for {
				c <- <-chCopy
			}
		}()
	}
	// 开启了多个goroutine，谁先有数据，就把数据送给c
	return c
}

// 也可以利用go语言中，值传递的特性
func fanIn2(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		// ch变量只有一份，所以后面的channel会覆盖前面的，所以这里需要拷贝一份传过去
		go func(in chan string) {
			// 将c1的数据取出送给c
			for {
				c <- <-in
			}
		}(ch)
	}
	// 开启了多个goroutine，谁先有数据，就把数据送给c
	return c
}

// 非阻塞等待
func nonBlockingWait(ch chan string) (string, bool) {
	select {
	case m := <-ch:
		return m, true
	default:
		return "", false
	}
}

// 超时等待
func timeoutWait(ch chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-ch:
		return m, true
		// timeout后会发送一个信号
	case <-time.After(timeout):
		return "timeout", false
	}
}

func main() {
	done := make(chan struct{})
	m1 := msgGen2("service1", done)
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout for service1")
		}
	}
	// struct{}:定义一个空的struct
	// struct{}{}:定义并初始化
	done <- struct{}{}
	// 阻塞获取任务结束的信号
	<-done
}
