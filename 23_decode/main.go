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
	jsonDataFromWeb := []byte(`{
		"coursename" : "ReactJS" ,
		"price" : 299, 
		"website" :"hobbs server", 
		"password" : "123", 
		"tags" : ["web-dev", "js"]
	}`)

	var lcoCourse course

	validation := json.Valid(jsonDataFromWeb)

	if validation {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData = make(map[string]interface{})

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Println("This is map", myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}
