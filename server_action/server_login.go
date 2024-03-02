package server_action

import (
	"log"
	"net/http"
	"go_websocket_server/server_core"
	"github.com/gorilla/websocket"
)

func ModCallTest() {
	log.Println("Login Mod Call Test")
}

var upgrader = websocket.Upgrader{} // use default options

func Login(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		// Write Login Logic . . .
	}
}

