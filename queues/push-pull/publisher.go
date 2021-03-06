package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

const serverAddr = "localhost:61613"
const queueName = "/queue/go_test"

const messageCount = 10

func sendMessages() {
	conn, err := stomp.Dial("tcp", serverAddr, nil)
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}
	println("connected to server", serverAddr)

	defer conn.Disconnect()

	for i := 1; i <= messageCount; i++ {
		text := fmt.Sprintf("Message #%d", i)
		err = conn.Send(queueName, "text/plain", []byte(text), nil)
		if err != nil {
			println("failed to send to server", err)
			return
		}
		println("message sent")
	}
	println("sending EXIT message")
	err = conn.Send(queueName, "text/plain", []byte("EXIT"), nil)
	if err != nil {
		println("failed to send EXIT message to server", err)
		return
	}
	println("message sent")
	println("sender finished")
}

func main() {
	sendMessages()
}
