package main

import (
	"io"
	"log"
	"net"
	"os"
)

/*
聊天客户端代码
*/
func main() {
	ChatClient()
}

func ChatClient() {
	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
