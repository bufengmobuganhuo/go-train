package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		// UTF-8，中文3字节，英文1字节
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// ch的类型是一个rune，4字节
	for i, ch := range s {
		// 会转化成unicode编码输出
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	// 获取到字符长度（=9）
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		// 获取到单个字符和对应的byte长度
		ch, size := utf8.DecodeRune(bytes)
		// 截取当前ch的长度
		bytes = bytes[size:]
		fmt.Printf("(%d %c)", size, ch)
	}
	fmt.Println()

	// 获取每个字符
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}
