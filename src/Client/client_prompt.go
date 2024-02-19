package client

import (
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

func (prompt *clientPrompt) WriteMessage(m string) {
	log.Println("Write Message")
}
