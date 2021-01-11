package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// 채널에 넣기
	ch <- 1
	ch <- 2

	// 채널 닫기
	close(ch)

	// 방법1
	// 무한 루프
	// for {
	// 	if i, success := <-ch, !success{
	// 		fmt.Println(i)
	// 	} else {
	// 		break
	// 	}
	// }

	// 방법2
	// for range
	for i := range ch {
		fmt.Println(i)
	}
}
