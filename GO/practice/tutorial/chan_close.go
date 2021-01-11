package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// 채널로 보내기
	ch <- 1
	ch <- 2

	// 채널 닫기
	close(ch)

	// 채널에서 받기
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	if _, success := <-ch; !success {
		fmt.Println("no more data")
	}
}
