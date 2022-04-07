package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	// 并发度10
	for i := 0; i < 10; i++ {
		go func(ii int) {
			// 一个死循环
			for {
				a[ii]++
				// 手动交出控制权，因为最新版已经支持了抢占式多任务处理，所以其实并不需要手动调用
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
