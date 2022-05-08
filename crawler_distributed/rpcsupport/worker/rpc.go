package worker

import "mengyu.com/gotrain/crawler/engine"

type CrawlService struct{}

// 封装执行爬虫的worker方法
func (CrawlService) Process(req Request, result *ParseResult) error {
	// 接收到的request是一个序列化的，这里反序列化
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	// 调用worker
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return nil
	}
	// 序列化result
	*result = SerializeResult(engineResult)
	return nil
}
