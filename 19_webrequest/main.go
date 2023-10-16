package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Welcome to Web Requests")

	resp, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", resp)

	defer resp.Body.Close() // callers respsibility to close the response connection

	databyte, err := io.ReadAll(resp.Body)

	checkNilErr(err)

	fmt.Println("Response data is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
