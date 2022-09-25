package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?course=reactjs&paymentid=123abcd"

func main() {

	fmt.Println("URLs in golang")

	fmt.Println(myUrl)

	// PARSING THE URL
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type queryParams: %T\n", queryParams)

	fmt.Println("queryParams: ", queryParams["course"])

	for _, value := range queryParams {
		fmt.Println("param value is ", value)

	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/css",
		RawPath: "user=gaurav",
	}
	fmt.Println("partsOfUrl ", partsOfUrl)
	fmt.Println("new URL ", partsOfUrl.String())

}
