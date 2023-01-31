package main

import (
	"chat/messages"
	"fmt"
	"log"
	"net"
	"os"
)

func handleClient(msgHandler *messages.MessageHandler) {
	defer msgHandler.Close()

	for {
		wrapper, _ := msgHandler.Receive()

		switch msg := wrapper.Msg.(type) {
		case *messages.Wrapper_RegistrationMessage:
			fmt.Println("Got a registration message. Not implemented yet!")
		case *messages.Wrapper_ChatMessage:
			fmt.Println("<"+msg.ChatMessage.GetUsername()+"> ",
				msg.ChatMessage.MessageBody)
		case nil:
			log.Println("Received an empty message, terminating client")
			return
		default:
			log.Printf("Unexpected message type: %T", msg)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			msgHandler := messages.NewMessageHandler(conn)
			// only handles one client at a time:
			handleClient(msgHandler)
		}
	}
}
