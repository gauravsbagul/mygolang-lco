package main

import "fmt"

func main() {
	var userName string = "Gaurav"
	fmt.Println("userName: ",userName)
	fmt.Printf("Variable is of type: %T \n",userName)
	
	var isLoggedIn bool = true
	fmt.Println("isLoggedIn: ",isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)
	
	var smallValue uint8 = 255
	fmt.Println("smallValue: ",smallValue)
	fmt.Printf("Variable is of type: %T \n",smallValue)

	var smallFloat float32 = 255.45523456765432
	fmt.Println("smallFloat: ",smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)
	
	var bigFloat float64 = 255.45523456765432
	fmt.Println("bigFloat: ",bigFloat)
	fmt.Printf("Variable is of type: %T \n",bigFloat)
	
	var anotherVariable int
	fmt.Println("anotherVariable: ",anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	var anotherVariableStr string
	fmt.Println("anotherVariableStr: ",anotherVariableStr)
	fmt.Printf("Variable is of type: %T \n",anotherVariableStr)

}
