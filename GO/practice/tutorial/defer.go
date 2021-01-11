package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("ex1.txt")
	if err != nil { // 파일 열기에 실패할 경우
		panic(err) // 실행 중지, 바로 defer 함수들 실행 후 리턴
	}

	// main 마지막에 파일 close 실행
	// defer fmt.Println("2222")
	// defer fmt.Println("2222")
	defer funcCall()
	// defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
}

func openFile 
