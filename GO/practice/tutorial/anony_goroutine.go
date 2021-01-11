package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait group 생성, 2개의 go 루틴을 기다림
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("func 1 done")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println("func 2 done, msg : ", msg)
	}("2222") // 파라미터로 "2222" 전달함

	wait.Wait() // Go 루틴이 모두 끝날때까지 대기
}
