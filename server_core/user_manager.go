package server_core

import (
	"go_websocket_server/user_info"
)

type UserManager struct {
	container []user_info.User
}

func (um UserManager) AddNewUser(newUser user_info.User) {
	um.container = append(um.container, newUser)
}
