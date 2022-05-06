package main

import (
	"testing"
	"time"

	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/model"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// 启动服务
	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	expected := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var result string
	err = client.Call("ItemSaverService.Save", expected, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
