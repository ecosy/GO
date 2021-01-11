package main

import "fmt"

func main() {
	// 정수형 채널 생성
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123 보내기
	}()

	var i int
	i = <-ch // 채널로부터 123 받기
	fmt.Println(i)
}
