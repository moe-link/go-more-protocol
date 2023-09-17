package main

import (
	"go-more-protocol/config"
	"go-more-protocol/protocol/ws"
)

func main() {
	ws.Init()
	config.Test()
}
