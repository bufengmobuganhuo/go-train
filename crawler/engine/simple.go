package engine

import (
	"log"
)

type SimpleEngine struct{}

// 传入多个request类型，单任务版的engine
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	// 一直运行，直到requests为空
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(req)
		if err != nil {
			continue
		}
		// 新产生的request放入requests
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
