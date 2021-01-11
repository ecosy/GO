package main

import (
	"fmt"
)

func main() {
	strs := []string{"n1", "n2", "n3"}
	for i, val := range strs {
		fmt.Println(i, val)
	}
}
