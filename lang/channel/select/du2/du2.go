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

var verbose = flag.Bool("v", false, "show verbose progress messages")

// 通知所有goroutine，当前任务是否已取消
var done = make(chan bool)

func main() {
	flag.Parse()
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		// 读取一个byte，如果读取到了，则说明结束了
		os.Stdin.Read(make([]byte, 1))
		// 关闭done
		close(done)
	}()

	var wg sync.WaitGroup
	filesizes := make(chan int64)
	for _, dir := range roots {
		wg.Add(1)
		go walkDir(dir, &wg, filesizes)
	}

	// 等待程序运行结束
	go func() {
		wg.Wait()
		close(filesizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(5 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		// 如果channel关闭，则会读取到零值
		case <-done:
			// 把通道里的数据接收完
			for size := range filesizes {
				nfiles++
				nbytes += size
			}
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func canceled() bool {
	select {
	// channel关闭后，会读取到零值，说明结束了
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", time.Now().Format("2006-01-02 15:04:05.000"), nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, filesizes chan<- int64) {
	// 执行完后，计数器-1
	defer wg.Done()
	// 如果已经发了结束信号，则不执行任务
	if canceled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			go walkDir(filepath.Join(dir, entry.Name()), wg, filesizes)
		} else {
			filesizes <- entry.Size()
		}
	}
}

var tokens = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case tokens <- struct{}{}:
		break
	case <-done:
		return nil
	}

	// 执行完后，释放资源
	defer func() {
		<-tokens
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
