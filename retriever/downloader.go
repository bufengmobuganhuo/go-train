package main

import (
	"fmt"

	"mengyu.com/gotrain/retriever/mock"
)

func getRetriever() retriever {
	return mock.Retriever{}
}

// 定义接口
type retriever interface {
	Get(url string) string
}

func main() {
	var retriever retriever = getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
