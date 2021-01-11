package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{width: 10., height: 20.} // struct
	c := Circle{radius: 10}            // struct

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for i, s := range shapes {
		a := s.area()

		fmt.Println(reflect.TypeOf(s))
		fmt.Println(i, s)
		fmt.Println("s.area() func call : ", a)
	}
}
