package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers Study")

	//var one int = 1

	//var ptr *int

	//fmt.Println("Value of ptr is", ptr)
	myNumber := 23

	var myNumberPtr = &myNumber

	fmt.Println("Value of myNumber is", myNumber)        // direct value
	fmt.Println("Value of myNumberPtr is", myNumberPtr)  // reference to direct memory location
	fmt.Println("Value of myNumberPtr is", *myNumberPtr) // dereferencing

	*myNumberPtr = *myNumberPtr + 1

	fmt.Println("Value of new myNumber is", myNumber) // direct value updated
}
