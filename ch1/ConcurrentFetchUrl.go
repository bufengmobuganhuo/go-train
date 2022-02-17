package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/**
并发获取URL
*/
func main() {
	start := time.Now()
	// 创建一个传递string类型参数的channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go表示创建一个新的goroutine
		// goroutine是函数的一种并发执行方式，
		// channel用来在goroutine之间进行参数传递
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// "<-"负责从 channel中接收结果，然后打印出来
		// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel
		// 里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 把打印结果给ch
		ch <- fmt.Sprint(err)
		return
	}
	// 将响应内容拷贝到ioutil.Discard中（类似一个垃圾桶，存储不需要的数据）
	// 此处只需要字节数，而不需要内容
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 将一个字符串写入ch
	ch <- fmt.Sprintf("%s.2fs    %7d    %s", secs, nbytes, url)
}
