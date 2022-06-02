package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer

	if debug {
		buf = new(bytes.Buffer)
	}

	// 把一个还未初始化的变量赋值给一个接口
	f(buf)

	if debug {
		//
	}
}

func f(out io.Writer) {
	// 接口的动态类型不为空（为*bytes.Buffer），但是动态值是nil，所以(out != nil) = true
	// 但是他没有指向任何具体Buffer，那么调用.Write方法时程序会崩溃，因为Write方法要求接受者不可为空
	// 有些方法是允许空接收者的
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

