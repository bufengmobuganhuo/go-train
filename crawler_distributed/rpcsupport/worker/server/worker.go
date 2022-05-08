package main

import (
	"flag"
	"fmt"
	"log"

	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	// 从命令行读取输入的port
	// go run worker.go --help 输出如何设置port
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// 开启执行爬虫的RPC服务
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))

}
