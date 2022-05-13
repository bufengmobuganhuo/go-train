package main

import (
	"bufio"
	"fmt"
	"os"

	"mengyu.com/gotrain/pipeline/nodes"
)

func main() {
	const filename = "small.in"
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := nodes.RandomSource(n)
	writer := bufio.NewWriter(file)
	nodes.WriterSink(writer, p)
	writer.Flush()
	// file, err = os.Open(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// p = nodes.ReaderSource(file)

	// for v := range p {
	// 	fmt.Println(v)
	// }
}

func mergeDemo() {
	p := nodes.Merge(nodes.InMemSort(nodes.ArraySource(3, 2, 6, 7, 4)), nodes.InMemSort(nodes.ArraySource(-1, 2, 5, 7, 4)))
	for v := range p {
		fmt.Println(v)
	}
}
