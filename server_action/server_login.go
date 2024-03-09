package server_action

import (
	"log"
	"net/http"
)

// var upgrader = websocket.Upgrader{} // use default options

func Login(w http.ResponseWriter, r *http.Request) {
	c, err := GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		// Write Login Logic . . .
	}
}
