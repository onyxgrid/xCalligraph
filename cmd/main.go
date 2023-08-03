package main

import (
	"github.com/paulrouge/xcall-event-watcher/internal/icon"
)

func main() {

	go icon.CallExecuteCall()
	go icon.CheckBlocks()
	go icon.HandleBlock()
	go icon.HandleTransaction()
	select {}
}
