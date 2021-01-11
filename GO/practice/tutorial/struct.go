package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "first"
	p.age = 10
	fmt.Println(p)

	var p2 person
	p2 = person{"second", 20}
	fmt.Println(p2)

	p3 := person{name: "third", age: 30}
	fmt.Println(p3)

	p4 := new(person)
	p4.name = "4rd"
	p4.age = 40
	fmt.Println(p4)

}
