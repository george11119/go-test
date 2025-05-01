package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNums := make([]int, 100)
	for i := range 100 {
		randomNums[i] = rand.Intn(101)
	}

	fmt.Println(randomNums)

	for _, v := range randomNums {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

	// shadowed by total declaration in for loop
	var total int
	for i := range 10 {
		total := total + i
		fmt.Println(total)
	}
}
