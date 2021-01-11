package main

func main() {
	add := func(i int, j int) int {
		return i + j
	}

	// 1. add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 2. 직접 파라미터에 익명 함수 정의
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
