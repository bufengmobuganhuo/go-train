package main

import (
	"fmt"
	"testing"
	"time"

	"mengyu.com/gotrain/crawler_distributed/config"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	// 可序列化的request
	req := worker.Request{
		Url: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}
	result := &worker.ParseResult{}
	err = client.Call(config.CrawlServiceRpc, req, result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

}
