package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 打印客户端的响应
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	log.Println("done in server")
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	// 从客户端一直读取，当客户端关闭连接时，判定没有输入，for循环结束
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// 关闭连接
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
