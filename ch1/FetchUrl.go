package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/**
获取 URL
*/
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// body包括一个刻度的服务器响应流
		// ioutil.ReadAll会获取到全部内容并保存到变量b中
		b, err := ioutil.ReadAll(resp.Body)
		// 关闭流
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
