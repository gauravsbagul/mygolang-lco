package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	fmt.Println("Web request in golang")

	response, err := http.Get(myUrl)
	if err != err {
		panic(err)
	}

	fmt.Printf("Type of response %T\n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != err {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println("content", content)

	// CALL API

	fmt.Println("API CALLING")

	PerformGETRequest()
	PerformPOSTJsonRequest()
	PerformPostFormRequest()
}

func PerformGETRequest() {

	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("PerformGETRequest-> Status code:", response.StatusCode)
	fmt.Println("PerformGETRequest-> Content length code:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("PerformGETRequest-> Content ", string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("PerformGETRequest-> byteCount ", byteCount)

	fmt.Println("PerformGETRequest-> byteCount ", responseString.String())

}

func PerformPOSTJsonRequest() {

	const myUrl = "http://localhost:3000/post"

	// FAKE JSON PAYLOAD

	requestBody := strings.NewReader(`
		{
			"courseName":"golang",
			"price":10,
			"platform":"lco.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("PerformPOSTJsonRequest-> err ", err)
		panic(err)
	}
	fmt.Println("PerformPOSTJsonRequest-> Content ", string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("PerformPOSTJsonRequest-> byteCount ", byteCount)

	fmt.Println("PerformPOSTJsonRequest-> responseString ", responseString.String())

}

func PerformPostFormRequest() {

	const myUrl = "http://localhost:3000/postForm"

	// FAKE FORM DATA

	data := url.Values{}

	data.Add("firstName", "Gaurav")
	data.Add("lastName", "Bagul")
	data.Add("email", "Gaurav@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("PerformPostFormRequest-> err ", err)
		panic(err)
	}
	fmt.Println("PerformPostFormRequest-> Content ", string(content))

}
