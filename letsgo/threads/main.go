package main

import (
	"fmt"
	"time"

	"github.wdf.sap.corp/i332859/learn-go/letsgo/threads/thread"
)

func main() {
	ch := make(chan bool, 2)
	go thread.Service(ch)
	fmt.Println("World")

	time.Sleep(5 * time.Second)
	ch <- false

	time.Sleep(5 * time.Second)
	ch <- true

	time.Sleep(15 * time.Second)

}
