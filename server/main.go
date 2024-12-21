package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns  []net.Conn
	connCh = make(chan net.Conn)
	done   = make(chan net.Conn)
	msgCh  = make(chan string)
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}
	// listen from client
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("New connection from", conn.RemoteAddr())
			//push new connect from client
			conns = append(conns, conn)
			connCh <- conn

		}
	}()

	for {
		select {
		case conn := <-connCh:
			go onMessage(conn)
		case conn := <-done:
			fmt.Println("client exit")
			removeConn(conn)
		case msg := <-msgCh:

			fmt.Print(msg)
		}
	}
}
func onMessage(conn net.Conn) {

	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msgCh <- msg
		publishMsg(conn, msg)
	}

	done <- conn
}
func publishMsg(conn net.Conn, msg string) {
	for i := range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}

}
func removeConn(conn net.Conn) {
	for i := range conns {
		if conns[i] == conn {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
}
