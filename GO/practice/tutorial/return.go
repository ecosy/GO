package main

import "fmt"

func main() {
	total, count := sum(1, 2, 3, 4)
	fmt.Println(total, count)
}

func sum(nums ...int) (int, int) {
	s := 0
	var count = 0
	for _, n := range nums {
		s += n
		count++
	}
	return s, count
}
