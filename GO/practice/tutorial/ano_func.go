package main

func main() {
	sum := func(n ...int) int {
		s := 0
		for _, v := range n {
			s += v
		}
		return s
	}

	result := sum(1, 2, 3)
	println(result)
}
