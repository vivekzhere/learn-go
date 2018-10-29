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

func Service2(ch chan bool) {
	for {
		select {
		case val := <-ch:
			if val {
				fmt.Println("Im done")
				return
			}
			fmt.Println("Im not done")
		default:
			fmt.Println("Hello")
			time.Sleep(1 * time.Second)
		}

	}
}
