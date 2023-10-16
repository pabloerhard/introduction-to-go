package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays Study")

	var myArray [5]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[3] = 4

	fmt.Println("Value of myArray is", myArray)
	fmt.Println("Length of myArray is", len(myArray))

	var myArray2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Value of myArray2 is", myArray2)

}
