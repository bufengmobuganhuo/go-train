package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"mengyu.com/gotrain/funtional/fib"
)

// 定义一个类型，这个类型实际是一个入参为空，返回为int的函数
type intGen func() int

// 类型实现Reader接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	// 结束
	if next > 10000 {
		return 0, io.EOF
	}
	// 将next读入字符串s
	s := fmt.Sprintf("%d\n", next)
	// 将s写入p数组
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 一直读，直到读不到为止
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
