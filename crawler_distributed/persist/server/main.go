package main

import (
	"log"

	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler_distributed/persist"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
)

func main() {
	// 打印异常日志，并退出
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	// 此处需要传入地址，因为ItemSaverService.Save是一个指针
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
