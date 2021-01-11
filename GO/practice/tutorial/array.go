package main

import "fmt"

func main() {
	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i, v := range a {
		fmt.Println(i, v)

		for _j, _v := range a[i] {
			fmt.Println(_j, _v)
		}
		fmt.Println()
	}
	fmt.Println("length : ", len(a))
}
