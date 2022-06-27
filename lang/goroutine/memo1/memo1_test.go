package memo1_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"mengyu.com/gotrain/lang/goroutine/memo1"
)

func TestGet(t *testing.T) {
	incomingUrls := []string{
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}
	var wg sync.WaitGroup
	m := memo1.New(memo1.HttpGetBody)
	for _, url := range incomingUrls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Wait()
		}(url)
	}
	wg.Wait()
}
