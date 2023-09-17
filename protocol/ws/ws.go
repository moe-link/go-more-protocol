package ws

import (
	"github.com/gorilla/websocket"
)

func Init() {
	_ = websocket.Upgrader{}
}
