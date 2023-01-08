package main

import (
	"fmt"
	"tg-vk-bridge/tgfuncs"
	"time"
)

func main() {
	for {
		fmt.Println(tgfuncs.GetUpdates())
		time.Sleep(1 * time.Second)
	}
}
