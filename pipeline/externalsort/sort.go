package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"mengyu.com/gotrain/pipeline/netnodes"
	"mengyu.com/gotrain/pipeline/nodes"
)

const (
	output      = "../sorted.out"
	input_large = "../large.in"
	input_small = "../small.in"
)

func main() {
	sortByNetwork()
}

func sortByNetwork() {
	p := createNetworkPipeline(input_large, 400, 8)
	writeToFile(output, p)
	printFile(output)
}

func singleSort() {
	p := createPipeline(input_large, 800000000, 8)
	writeToFile(output, p)
	printFile(output)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	p := nodes.ReaderSource(reader, -1)
	count := 0
	for v := range p {
		count++
		fmt.Println(v)
		if count >= 100 {
			break
		}
	}
}

func writeToFile(filename string, p <-chan int) {
	// 如果文件存在，则删除
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	nodes.WriterSink(writer, p)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := (fileSize + chunkCount - 1) / chunkCount
	sortAddr := []string{}
	nodes.Init()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		// 从文件中分块读取数据
		source := nodes.ReaderSource(bufio.NewReader(file), chunkSize)
		// 在内存中排好序后，传输给网络节点
		addr := ":" + strconv.Itoa(9000+i)
		netnodes.NetworkSink(addr, nodes.InMemSort(source))
		// 收集排序好的channel
		sortAddr = append(sortAddr, addr)
	}
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		// 从网络中读取数据
		sortResults = append(sortResults, netnodes.NetworkSource(addr))
	}
	// 将多个channel merge成一个
	return nodes.MergeN(sortResults...)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := (fileSize + chunkCount - 1) / chunkCount
	sortResults := []<-chan int{}
	nodes.Init()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		// 从文件中分块读取数据
		source := nodes.ReaderSource(bufio.NewReader(file), chunkSize)
		// 在内存中排序
		sortedChan := nodes.InMemSort(source)
		// 收集排序好的channel
		sortResults = append(sortResults, sortedChan)
	}
	// 将多个channel merge成一个
	return nodes.MergeN(sortResults...)
}
