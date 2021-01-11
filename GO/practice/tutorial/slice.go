package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int
	a = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
