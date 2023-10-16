package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web")
	PerformPostRequest()
	PerformGetRequest()
}

// PerformGetRequest sends a GET request to the specified URL and prints the response status code, content length, and body.
func PerformGetRequest() {
	// Define the URL to send the GET request to.
	const myUrl string = "http://localhost:8080/group/getAllGroups"

	// Send the GET request and handle any errors.
	resp, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	// Print the response status code and content length.
	//fmt.Println("Status code is:", resp.StatusCode)
	//fmt.Println("Content length is:", resp.ContentLength)

	// Read the response body and print it.
	var responseString strings.Builder
	content, _ := io.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())

	// Close the response body.
	defer resp.Body.Close()
}

func PerformPostRequest() {
	const myUrl string = "http://localhost:8080/group/create"
	requestBody := strings.NewReader(`{
		"name" : "Tester",
		"area" : "This is a basic account"
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
