package main

import (
	"bufio"
	"chat/messages"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	user := os.Args[1]
	fmt.Println("Hello, " + user)

	host := os.Args[2]
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer conn.Close()

	msgHandler := messages.NewMessageHandler(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("message> ")
		result := scanner.Scan() // Reads up to a \n newline character
		if result == false {
			break
		}

		message := scanner.Text()
		if len(message) != 0 {
			msg := messages.Chat{Username: user, MessageBody: message}
			wrapper := &messages.Wrapper{
				Msg: &messages.Wrapper_ChatMessage{ChatMessage: &msg},
			}
			msgHandler.Send(wrapper)
		}
	}
}
