package main

/*
聊天服务端代码
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ChatServer()
}

func ChatServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // 对外发送消息的通道

var ( // 所有客户端
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			/*
				把所有接收到的消息广播给所有客户端
				发送消息通道
			*/
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli) // 客户端离开，关闭通道
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户端消息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "欢迎" + who
	messages <- who + " 上线"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " 下线"
	conn.Close()

}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
