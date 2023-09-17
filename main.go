package main

import (
	"go-more-protocol/config"
	"go-more-protocol/proto/ws"
)

func main() {
	ws.Init()
	config.Test()
}
