package main

import (
	"fmt"

	"github.com/troylelandshields/easychat"
	"github.com/troylelandshields/easyinput"
)

func main() {
	fmt.Println("What's the IP address of the chat server?")
	ipAddress := easyinput.TakeInput()
	if ipAddress == "" {
		fmt.Println("You have to give an IP address to join a chat room!")
		return
	}

	fmt.Println("Enter your name: ")
	name := easyinput.TakeInput()
	if name == "" {
		fmt.Println("You must give a name!")
		return
	}

	chatClient, err := easychat.JoinChatRoom(ipAddress, name)
	if err != nil {
		fmt.Println("Error occurred when joining chatroom", err.Error())
		return
	}

	go receiveMessagesLoop(chatClient)

	sendMessagesLoop(chatClient)
}

func receiveMessagesLoop(chatClient *easychat.ChatClient) {
	for {
		msg, ok := chatClient.ReceiveMessage()
		if !ok {
			return
		}

		fmt.Printf("\n[%s] %s: %s\n...> ",
			msg.Time.Format("Jan 2, 3:04:05 PM"), msg.From, msg.Body)
	}
}

func sendMessagesLoop(chatClient *easychat.ChatClient) {
	for {
		fmt.Print("> ")
		msg := easyinput.TakeInput()

		if msg == "" {
			continue
		}

		chatClient.SendMessage(msg)
	}
}
