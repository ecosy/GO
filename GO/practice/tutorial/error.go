package main

import (
	"fmt"
	"log"
)

func main() {
	// f, err := os.Open("./ex.txt")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println(f.Name())

	_, err = otherFunc()

	switch err.(type) {
	default: // no error
		fmt.Println("no error")
	case MyError:
		log.Print("log my error")
	case error:
		log.Fatal(err.Error())
	}
}
