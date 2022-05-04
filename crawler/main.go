package main

import (
	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/persist"
	"mengyu.com/gotrain/crawler/scheduler"
	"mengyu.com/gotrain/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver()

	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.SimpleScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	// 启动入口入口
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
