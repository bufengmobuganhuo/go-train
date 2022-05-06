package main

import (
	"fmt"
	"time"

	"mengyu.com/gotrain/retriever/mock"
	"mengyu.com/gotrain/retriever/real"
)

func getRetriever() Retriever {
	return &real.Retriever{}
}

// 定义接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

// 函数传入一个接口
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	// type switch
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var retriever Retriever = getRetriever()
	fmt.Println(retriever.Get(url))
	var r Retriever
	r = &mock.Retriever{Contents: "this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	// type assertion，类似java的强转
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	}

	s := mock.Retriever{Contents: "this is a fake imooc.com"}
	fmt.Println("Try a session")
	fmt.Println(session(&s))
}
