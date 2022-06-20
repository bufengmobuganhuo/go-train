package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// 使用缓冲通道，限制并发请求在20个以内
var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	// 发送到任务列表的数量
	var n int

	// 从命令接收第一个要爬取的url
	n++
	// 一开始没有goroutine接收，会一直阻塞
	go func() {
		worklist <- os.Args[1:]
	}()

	// 去重
	seen := make(map[string]bool)
	// 一直执行，直到发送到任务列表的数量=0
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				// 发送到任务列表的任务数量+1
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	// 获取令牌，如果缓冲空间用完，会阻塞，直到有其他goroutine发送令牌到tokens
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// 释放令牌
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
