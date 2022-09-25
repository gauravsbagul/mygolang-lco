package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is maps:")

	lang := make(map[string]string)

	lang["js"] = "java script"
	lang["rb"] = "ruby"
	lang["py"] = "python"

	fmt.Println("list of all: ", lang)
	fmt.Println("js is : ", lang["js"])

	delete(lang, "rb")
	fmt.Println("list of all: ", lang)

	// loops

	for key, value := range lang {
		fmt.Printf("Key-> %v, value-> %v\n", key, value)

	}
}
