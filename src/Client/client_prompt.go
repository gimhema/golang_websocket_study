package client

import (
	"fmt"
	"log"
)

type clientPrompt struct {
	pid int
}

func createClientPrompt() *clientPrompt {
	_prompt := clientPrompt{pid: -1}

	return &_prompt
}

func (prompt *clientPrompt) Init() {
	log.Println("Prompt User Init")
}

func (prompt *clientPrompt) WriteMessage() string {
	var message = ""
	fmt.Print("Write Message : ")
	fmt.Scanln(&message)

	return message
}
