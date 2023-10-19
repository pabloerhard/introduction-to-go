package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB Api")
	fmt.Println("Getting Started")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
