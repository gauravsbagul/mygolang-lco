package main

import "fmt"

func main() {

	fmt.Println("Strut in golang")

	gaurav := User{"Gaurav", "g@gmail.com", true, 29}

	fmt.Println("User", gaurav)

	fmt.Printf("User details are %+v\n", gaurav)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
