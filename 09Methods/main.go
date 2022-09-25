package main

import "fmt"

func main() {

	fmt.Println("Strut in golang")

	gaurav := User{"Gaurav", "g@gmail.com", true, 29}

	fmt.Println("User", gaurav)

	fmt.Printf("User details are %+v\n", gaurav)

	gaurav.GetStatus()

	gaurav.NewMail()

	fmt.Printf("User details are %+v\n", gaurav)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active?", u.Status)
}

func (u User) NewMail() {

	u.Email = "test@go.dev"
	fmt.Println("Is user Active?", u.Email)
}
