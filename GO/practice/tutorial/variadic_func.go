package main

import "fmt"

func main() {
	say("1", "2", "3")
	say("4")
}

func say(msg ...string) {
	for i, v := range msg {
		fmt.Println(i, v)
	}
}
