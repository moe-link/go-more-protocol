package main

import (
	"fmt"
	"go-more-protocol/config"
	"go-more-protocol/proto/ws"
	"os"
	"os/signal"
	"time"
)

func init() {
	ws.Init()
	config.Test()
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		fmt.Println("Go routine running")
		time.Sleep(3 * time.Second)
		fmt.Println("Go routine done")
	}()
	<-c
	fmt.Println("bye")
}
