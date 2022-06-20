package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 用户之间发送消息，同时
type client chan<- string

var (
	// 用于通知有新用户加入
	entering = make(chan client)
	// 用于接收所有用户发送的消息
	messages = make(chan string)
	// 用于通知用户离开
	leaving = make(chan client)
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
			log.Printf("arrived")
		case msg := <-messages:
			// 把消息发送给所有用户
			for cli := range clients {
				cli <- msg
			}
		case cli := <-leaving:
			delete(clients, cli)
			// 关闭用户发送消息的通道
			close(cli)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// 把用户对应的channel里的消息都写入到网络中
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleConn(conn net.Conn) {
	// 用户发送消息对应的channel
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	// 告诉用户，给它分配的用户名
	ch <- "You are " + who
	// 广播所有用户，告知有新用户到来
	messages <- who + " has arrived"
	entering <- ch

	// 从网络中读取用户输入的数据
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// 把用户输入的消息广播给所有人
		messages <- who + ": " + scanner.Text()
	}
	// 用户离开，通知所有人
	leaving <- ch
	messages <- who + "has left"
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
