package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

/**
Web服务器
*/
func main() {
	// 将处理函数和URL连接起来,
	// 对于每个传入的请求，服务器会在不同的goroutine中运行处理函数
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/request", requestInfo)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 写入响应
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

// 并发计数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// 回显HTTP请求
// r *http.Request：获取指针引用的变量的值
func requestInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// 先赋值，再判断
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
