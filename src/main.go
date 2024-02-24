package main

import (
    "go_websocket_server/sub_module"
)

func main() {
    for _, e := range sub_module.SubModuleCall() {
        println(e)
    }
}