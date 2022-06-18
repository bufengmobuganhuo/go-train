package netnodes

import (
	"bufio"
	"net"

	"mengyu.com/gotrain/pipeline/nodes"
)

// 网络连接
func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		// 写入文件
		nodes.WriterSink(writer, in)
	}()

}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		// 从网络中读取数据
		r := nodes.ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
