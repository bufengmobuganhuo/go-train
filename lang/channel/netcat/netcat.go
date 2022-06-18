package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		// 将服务端的回显输出到控制台
		// 会一直读取，直到遇到EOF，或者错误
		// 服务端关闭连接后，Copy方法结束
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// 告知主goutine结束
		done <- struct{}{}
	}()
	// 接收控制台的输入，并且拷贝到connection中，也就是发送给服务端
	mustCopy(conn, os.Stdin)
	log.Println("received in stdin")
	// mustCopy方法结束后，关闭连接（此时会导致服务端关闭连接，最终导致第20行代码对应的goroutine继续向下执行）
	conn.Close()
	// 在这个channel，会先执行这里的接收操作，此时主gorouine被阻塞，直到有一个值发送到channel中
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	// 输入(control + D)后，服务端会认为客户端接收完毕，此时会关闭连接，那么Copy方法结束
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
