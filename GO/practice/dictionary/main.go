package main

import (
	"fmt"

	mydict "./dict"
)

func main() {
	// 1. Create Dictionary
	dictionary := mydict.Dictionary{
		"first": "First word",
	}

	// Search a word
	definition, err := dictionary.Search("first")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	switch err {
	case nil:
		fmt.Println(err)
	default:
		fmt.Println(definition)
	}

	fmt.Println(dictionary["first"])

	// 2. Add value
	err = dictionary.Add("second", "2 def")
	if err != nil {
		fmt.Println(dictionary)
	}

	for _key, _value := range dictionary {
		fmt.Println(_key, _value)
	}

	// 3. Update value
	err = dictionary.Update("second", "new definition")
	if err == nil {
		fmt.Println("updated key, value : ", "second", dictionary["second"])
	}
	err = dictionary.Update("3rd", "new definition")
	if err == nil {
		fmt.Println("updated key, value : ", "second", dictionary["second"])
	} else {
		fmt.Println(err)
	}

	// 4. Delete word
	err = dictionary.Delete("second")
	if err == nil {
		fmt.Println("delete completed!")
		fmt.Println(dictionary)
	} else {
		fmt.Println(err)
	}
}
