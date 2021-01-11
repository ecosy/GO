package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m map[int]string
	var m2 map[int][]int

	m = make(map[int]string)
	m2 = make(map[int][]int)

	m2[1] = []int{1, 2, 3}

	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	fmt.Println("m[1] : ", m[1])
	fmt.Println("m[-1] : ", m[-1])

	fmt.Println("m2[-1] : ", m2[-1] == nil)

	fmt.Println("type m[-1] : ", reflect.TypeOf(m[-1]))
	fmt.Println("type m2[-1] : ", reflect.TypeOf(m2[-1]))

	delete(m, 1)
	delete(m, 0)
}
