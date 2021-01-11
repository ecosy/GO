package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

// 송신 채널- 보내기
func sendChan(ch chan<- string) {
	ch <- "Data"
}

// 수신 채널 - 받기
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
