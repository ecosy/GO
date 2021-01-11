package main

func main() {
	var s []int

	if s == nil {
		println("it is nil ~~~")
	}
	println(len(s), cap(s))
}
