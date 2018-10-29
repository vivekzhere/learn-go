package thread

import (
	"fmt"
	"time"
)

func Service(ch chan bool) {
	done := false
	go func() {
		for {
			select {
			case val := <-ch:
				done = val
			}
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello")
		if done {
			fmt.Println("In done")
			return
		}

	}
}
