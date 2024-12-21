package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Print("Enter your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, err := nameReader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading name:", err)
	}
	nameInput = strings.TrimSpace(nameInput)

	go onMessage(conn)

	for {
		fmt.Print("Message: ")
		messageReader := bufio.NewReader(os.Stdin)
		msg, err := messageReader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		msg = strings.TrimSpace(msg)
		formattedMsg := fmt.Sprintf("%s: %s\n", nameInput, msg)
		conn.Write([]byte(formattedMsg))
	}
}

func onMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Disconnected from server.")
			return
		}
		fmt.Print(msg)
	}
}
