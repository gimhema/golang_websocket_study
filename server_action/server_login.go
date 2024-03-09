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
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("Login Message : %s,", message)

		// message를 분리해서 로그인 정보를 추출

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
