package engine

import (
	"log"

	"mengyu.com/gotrain/crawler/fetcher"
)

func Worker(req Request) (ParseResult, error) {
	log.Printf("Fetching %s", req.Url)
	// 根据URL获取页面
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", req.Url, err)
		return ParseResult{}, err
	}
	// 解析页面，返回结果
	return req.ParserFunc(body), nil
}
