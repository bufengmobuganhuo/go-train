package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/scheduler"
	"mengyu.com/gotrain/crawler/zhenai/parser"
	"mengyu.com/gotrain/crawler_distributed/config"
	itemsaver "mengyu.com/gotrain/crawler_distributed/persist/client"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
	workerclient "mengyu.com/gotrain/crawler_distributed/rpcsupport/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

// 分布式版
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	// 创建处理爬虫的worker池
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := workerclient.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	// 启动入口入口
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("error connecting to %s: %v", host, err)
			continue
		}
		log.Printf("connected to %s", host)
		clients = append(clients, client)
	}
	out := make(chan *rpc.Client)
	go func() {
		// 一直向channel发送创建好的client
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
