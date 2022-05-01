package main

import (
	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/zhenai/parser"
)

func main() {
	// 启动入口入口
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
