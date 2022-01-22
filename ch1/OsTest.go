package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//for1()
	//for2()
	//for3()
	//for4()
	for5()
}

func for1() {
	// 字符串会被自动初始化成""
	var s, sep string
	// i := 0 是短变量声明的一部分
	// --i这种形式是非法的
	for i := 0; i < len(os.Args); i++ {
		// 连接字符串
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Print(s)
}

func for2() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Print(s)
}

func for3() {
	var s, sep string
	i := 0
	// 如果连条件语句都省略，则是一个无限循环语句，此时可以使用break,return结束循环
	for i < len(os.Args) {
		s += sep + os.Args[i]
		sep = " "
		i++
	}
	fmt.Print(s)
}

func for4() {
	s, sep := "", ""
	// range关键字会产生一对"索引, 元素"，此处因为不需要索引，所以使用"_"替代
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Print(s)
}

func for5() {
	// 每个元素之间添加一个" "
	fmt.Print(strings.Join(os.Args[1:], " "))
}
