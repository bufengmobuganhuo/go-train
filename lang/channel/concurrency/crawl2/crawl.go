package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	// 未去重的urls
	worklist := make(chan []string)
	// 去重后的url
	unseenLinks := make(chan string)

	// 接收初始url
	go func() {
		worklist <- os.Args[1:]
	}()

	// 创建20个活跃的goroutine执行爬虫任务
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// 将发现的新url放到worklist
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	// 主goroutine中从worklist获取待爬取的任务，去重后放入unseenLinks
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
