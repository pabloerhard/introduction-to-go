package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices Study")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") // add to slice

	fruitList = append(fruitList[1:]) // remove first element

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 777) // this allocates new memory

	fmt.Println(highScores)
	sort.Ints(highScores) // sort slice
	fmt.Println(highScores)

}
