package main

import (
	"fmt"
	"time"

	"github.wdf.sap.corp/i332859/learn-go/letsgo/threads/thread"
)

func main() {
	ch := make(chan bool, 2)
	go thread.Service2(ch)

	time.Sleep(5 * time.Second)
	ch <- false

	time.Sleep(5 * time.Second)
	ch <- true

	time.Sleep(2 * time.Second)
	fmt.Println("Main done")
}
