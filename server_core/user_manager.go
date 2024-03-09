package server_core

import (
	"go_websocket_server/user_info"
)

type UserManager struct {
	user_container  []user_info.User
	user_search_map map[string]*user_info.User // Key : UserID
}

func (um UserManager) Initialize() {
	um.user_search_map = make(map[string]*user_info.User)
}

func (um UserManager) AddNewUser(newUser user_info.User) {
	um.user_container = append(um.user_container, newUser)
	um.user_search_map[newUser.UserID] = &newUser

}
