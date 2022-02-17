package main

import (
	"flag"
	"fmt"
	"strings"
)

// 创建一个新的bool标识变量，三个参数为：标识的名字，默认值，用户提供非法标识时的提示
// 返回的是一个指向标识变量的指针
var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

func main() {
	// 在使用标识前，必须调用flag.Parse来更新标识变量的默认值
	flag.Parse()
	// flag.Args获取到非标识参数，是一个slice
	// strings.Join，把slice使用*sep连接成一个字符串
	fmt.Print(strings.Join(flag.Args(), *sep))
	// *n = true, 忽略正常输出时结尾的换行符
	if !*n {
		fmt.Println()
	}
}
