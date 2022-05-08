package main

import (
	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/persist"
	"mengyu.com/gotrain/crawler/scheduler"
	"mengyu.com/gotrain/crawler/zhenai/parser"
	"mengyu.com/gotrain/crawler_distributed/config"
)

func main() {

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		// 单机版
		RequestProcessor: engine.Worker,
	}
	// 启动入口入口
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
