package main

import "fmt"

func main() {
	src := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}

	fmt.Println(src)

	target := make([][]int, len(src), cap(src))
	copy(target, src)

	fmt.Println(target)

	fmt.Println("&target[0] : ", &target[0])
	fmt.Println("&target : ", target)

	fmt.Println("&src[0] : ", &src[0])
	fmt.Println("&src : ", src)
}
