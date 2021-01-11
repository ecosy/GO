package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("func finished")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	length, up := lenAndUpper("My name")
	fmt.Println(length, up)
}
