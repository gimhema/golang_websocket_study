package main

import (
	"go_websocket_server/sub_module"

	"go_websocket_server/server_action"
)

func main() {
	for _, e := range sub_module.SubModuleCall() {
		println(e)
	}

	server_action.ModCallTest()

}
