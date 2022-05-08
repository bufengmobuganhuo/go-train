package client

import (
	"net/rpc"

	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler_distributed/config"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	// 将处理request的封装成调用RPC的
	return func(r engine.Request) (engine.ParseResult, error) {
		// 序列化request
		seriableReq := worker.SerializeRequest(r)
		var seriableRes worker.ParseResult
		// 从channel中取出一个client
		client := <-clientChan
		// RPC调用
		err := client.Call(config.CrawlServiceRpc, seriableReq, &seriableRes)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 反序列化
		return worker.DeserializeResult(seriableRes)
	}
}
