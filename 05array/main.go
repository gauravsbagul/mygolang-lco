package main

import "fmt"


func main(){

	fmt.Println("Welcome to array golang")

	var list [4]string

	list[0] = "Apple"
	list[1] = "orange"
	list[3] = "banana"

	fmt.Println("List is : ", list)
	fmt.Println("Length of the List is : ", len(list))

	var list2 = [3]string{"potato","beans","mushrooms"}

	fmt.Println("Veg list is : ", list2)
}

