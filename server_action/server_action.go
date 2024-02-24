package server_action

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ModCallTest() {
	log.Println("Mod Call Test")
}

var upgrader = websocket.Upgrader{} // use default options

func Echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
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
