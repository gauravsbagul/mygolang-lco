package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("This is JSON class in golang")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{
			"ReactJS BootCamp",
			299,
			"lco.dev",
			"abcd123",
			[]string{"web-dev", "js"},
		},
		{
			"MERN BootCamp",
			199,
			"lco.dev",
			"xyz1234",
			[]string{"fullstack", "js"},
		},
		{
			"Angular BootCamp",
			322,
			"lco.dev",
			"123bcd",
			nil,
		},
	}

	// PACKAGE THIS DATA AS JSON DATA

	finalJson, err := json.MarshalIndent(lcoCourses, "-", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON data %s\n", finalJson)

}

func DecodeJson() {

	jsonData := []byte(`
		{
            "courseName": "ReactJS BootCamp",
            "price": 299,
            "website": "lco.dev",
            "tags": ["web-dev", "js"]
       	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is Valid")

		json.Unmarshal(jsonData, &lcoCourse)

		fmt.Printf("lcoCourse JSON data %#v\n", lcoCourse)
	} else {
		fmt.Printf("JSON IS NOT VALID")
	}

	// EXTRACT KEY-VALUES

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonData, &myOnlineData)

	fmt.Printf("myOnlineData JSON data %#v\n\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T \n", key, value, value)
	}

}
