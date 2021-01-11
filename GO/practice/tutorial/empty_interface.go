package main

import "fmt"

func main() {
	var x interface{}
	x = 1
	x = "Tom"

	prnt(x)
}

func prnt(v interface{}) {
	fmt.Println(v)
}
