package server_action

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	once     sync.Once
	upgrader websocket.Upgrader
)

func GetUpgrader() *websocket.Upgrader {
	once.Do(func() {
	})
	return &upgrader
}

func Echo(w http.ResponseWriter, r *http.Request) {
	c, err := GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s,", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
