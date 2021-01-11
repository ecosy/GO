package main

import "fmt"

func main() {
	var a interface{} = 1

	i := a
	j1 := a.(int)

	fmt.Println("i : ", i)
	fmt.Println("j1 : ", j1)
}
