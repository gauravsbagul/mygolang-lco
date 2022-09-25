package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is slice:")

	var list = []string{"One", "Two", "Three"}

	fmt.Printf("Type of list is %T\n", list)
	list = append(list, "Four")

	fmt.Println("Type of list is : ", list)

	list = append(list[1:3])

	fmt.Println("Type of list is : ", list)

	score := make([]int, 4)

	score[0] = 234
	score[1] = 999
	score[2] = 878
	score[3] = 988

	score = append(score, 555, 123, 543)

	fmt.Println("score is : ", score)

	sort.Ints(score)
	fmt.Println("Sorted score : ", score)
	fmt.Println(sort.IntsAreSorted(score))

	// REMOVE VALUE BASED ON INDEX

	var index int = 3

	score = append(score[:index], score[index+1:]...)
	fmt.Println("Sorted score : ", score)

}
