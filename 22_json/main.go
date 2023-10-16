package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"password"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in Go")
	encodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "hobbs server", "123", []string{"web-dev", "js"}},
		{"MERN", 199, "hobbs server", "123", []string{"fullstack", "js"}},
		{"Angular", 199, "hobbs server", "123", nil},
	}

	// package json

	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	indentJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", indentJson)
}
