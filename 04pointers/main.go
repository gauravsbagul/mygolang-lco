package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	number := 23

	var myPtr = &number

	fmt.Println("Value of pointer is: ", myPtr)

	*myPtr += 40;
	fmt.Println("New Value of pointer is: ", *myPtr)

	fmt.Println("New Value of number is: ", number)


}