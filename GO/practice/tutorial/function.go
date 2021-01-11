package main

import (
	"fmt"
)

func main() {
	msg := "hi"
	say(msg)
}

func say(msg string) {
	fmt.Println(msg)
}
