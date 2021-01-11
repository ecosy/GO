package main

import (
	"fmt"
	"reflect"
)

func main() {
	map_ex := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	// key check
	val_1, check_1 := map_ex[1]
	val_0, check_0 := map_ex[0]
	fmt.Println("val, check 0 : ", val_0, check_0)
	fmt.Println("val 0's type : ", reflect.TypeOf(val_0))
	fmt.Println("val 0 == blank ? : ", val_0 == "")
	fmt.Println("val, check 0 : ", val_1, check_1)

	for k, v := range map_ex {
		fmt.Println(k, v)
	}
}
