package engine

import (
	"log"

	"mengyu.com/gotrain/crawler/fetcher"
)

// 传入多个request类型
func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	// 一直运行，直到requests为空
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", req.Url)
		// 根据URL获取页面
		body, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", req.Url, err)
			continue
		}
		// 将页面解析，此时可能产生新的request
		parseResult := req.ParserFunc(body)
		// 新产生的request放入requests
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}
