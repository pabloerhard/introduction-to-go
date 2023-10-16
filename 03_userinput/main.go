package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")

	// comma ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", input)
}
