package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")
	start := time.Now()
	fmt.Println("Current time is", start)
	fmt.Println(start.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2003, time.December, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created Date is", createdDate)
	fmt.Println("Created Date is", createdDate.Format("01-02-2006 Monday"))
}
