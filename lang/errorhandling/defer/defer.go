package main

import (
	"bufio"
	"fmt"
	"os"

	"mengyu.com/gotrain/lang/basic/func/funtional/fib"
)

func tryDefer() {
	// 函数退出是才会打印, 是一个栈，先进后出
	defer fmt.Println(1)
	defer fmt.Println(3)
	panic("error occurred")
}

func writeFile(filename string) {
	// 打开一个文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 将缓冲区的内容写入文件，defer就近写就可以
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
