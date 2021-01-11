package main

func main() {
	grade(50)
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	default:
		println("F")
	}
}

// A
