package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler_distributed/persist"
	"mengyu.com/gotrain/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// 打印异常日志，并退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), "dating_profile"))
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
