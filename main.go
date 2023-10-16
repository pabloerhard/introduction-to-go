package main

import "fmt"

const LoginToken string = "asdasdasd" // Capital letter means public, lowercase means private

func main() {
	var username string = "pabloerhard"
	var isLoggedIn bool = false

	fmt.Printf("Variable is of type: %T \n", username)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// no var style

	number := 42
	fmt.Println(number)

	fmt.Println(LoginToken)
}
