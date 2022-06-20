package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 并发地遍历目录，输出目录下文件的个数，和目录中文件的总大小

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(5 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", time.Now().Format("2022-09-09 12:22:12.987"), nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	// 执行完毕后给wg -1
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 令牌桶
var tokens = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	// 如果通道满了，会阻塞
	tokens <- struct{}{}
	// 释放令牌
	defer func() {
		<-tokens
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
