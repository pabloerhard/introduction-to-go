package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to My Files")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, error := os.Create("./myFile.txt")

	checkNilErr(error)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("Text data inside file is: ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
