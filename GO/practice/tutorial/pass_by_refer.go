package main

import "fmt"

func main() {
	msg := "origin"
	say(&msg)
	fmt.Println(msg)
}

func say(msg *string) {
	fmt.Println(*msg)
	*msg = "changed...."
}
