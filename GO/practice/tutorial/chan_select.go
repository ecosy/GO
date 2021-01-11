package main

import (
	"fmt"
	"time"
)

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			fmt.Println("run1 finished")
		case <-done2:
			fmt.Println("run2 finished")
			break EXIT
		default:
			fmt.Println("i am default")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("1 sec sleep")
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("2 sec sleep")
	done <- true
}
