package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
查找重复的行
*/
func main() {
	//findDuplicateLine()
	//findDuplicateLine2()
	findDuplicateLine3()
}

func findDuplicateLine3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// 返回的是一个字节切片，必须把它转化成string
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3:  %v\n", err)
			continue
		}
		// 转化成string，然后用换行符分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/**
查找给定文件内的重复行
*/
func findDuplicateLine2() {
	counts := make(map[string]int)
	// 读取命令行参数中的文件路径
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// 遍历文件路径
		for _, arg := range files {
			// 返回被打开的文件和可能的error
			f, err := os.Open(string(arg))
			if err != nil {
				_, _ = fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// 关闭资源
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/**

 */
func countLines(f *os.File, counts map[string]int) {
	// 读取文件的每一行
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

/**
查找控制台输入重复的行
*/
func findDuplicateLine() {
	// 创建一个Map，键是string类型，值是int类型
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// 读入下一行，并移除行末的换行符；如果读取到返回true
	for input.Scan() {
		line := input.Text()
		// 如果读取到空行，则结束
		if len(line) == 0 {
			break
		}
		// 获取读取的内容
		counts[input.Text()]++
	}

	// 迭代时,会迭代key-value
	for key, value := range counts {
		if value > 1 {
			/**
			%d			十进制整数
			%x,%o,%b	十六进制，八进制，二进制整数
			%f,%g,%e	浮点数：3.141593 3.141592653589793 3.141593e+00
			%t			布尔
			%c			字符
			%s			字符串
			%q			带双引号的字符串"abc"或单引号字符'a'
			%v			变量的自然形式
			%T			变量的类型
			%%			字面上的百分号
			*/
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
