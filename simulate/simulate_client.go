package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

const (
	messagePerClient = 10
	totalClient      = 1000
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < totalClient; i++ {
		wg.Add(1)
		go func(client int) {
			defer wg.Done()
			simulateClient(client)
		}(i)
	}
	wg.Wait()
	log.Print("all client done")
}
func simulateClient(ClientId int) {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("err while connecting to server: %v", err)
	}
	defer conn.Close()

	clientName := fmt.Sprintf("client%d", ClientId)
	var wg sync.WaitGroup
	for i := 0; i <= messagePerClient; i++ {
		wg.Add(1)
		go func(msgId int) {
			defer wg.Done()
			clientMessage := fmt.Sprintf("client message from %s: %d\n", clientName, msgId)
			_, err := conn.Write([]byte(clientMessage))
			if err != nil {
				log.Fatalf("err while sending message: %v", err)
			}
		}(i)

	}
	wg.Wait()
	log.Printf("finished sending message from client: %d", ClientId)
}
