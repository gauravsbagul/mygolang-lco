package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files in golang")

	content := "This needs to go in file - by gaurav"

	fmt.Println("File content", content)

	file, err := os.Create("./myGolangFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length", length)

	readFile("myGolangFile.txt")

	defer file.Close()
}

func readFile(fileName string) {

	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data in file is ", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}
