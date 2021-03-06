package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"mengyu.com/gotrain/crawler_distributed/config"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	// 等待100毫秒后再请求
	<-rateLimiter
	log.Printf("Fetching url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(reader)
}

// 确认html的字符集
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 读取1024字节
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
