package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	say("sync")

	go say("async 1")
	go say("async 2")
	go say("async 3")

	// 3초 대기
	time.Sleep(time.Second * 3)
}
