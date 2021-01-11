package main

import "fmt"

func main() {
	var a = 1
	for a < 15 {
		a += 1
		if a == 10 {
			goto END
		}
	}
END:
	fmt.Println("END")
}
