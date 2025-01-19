package main

import (
	"fmt"
)

func main() {

	// type one of for loop

	for i := 0; i < 10; i++ {

		fmt.Println(i)

	}

	// type two of for loop like while loop

	count := 0

	for count < 10 {
		fmt.Println("Count:", count)
		count++
	}

	// infinite loop

	for {
		fmt.Println("infinite loop")
		break
	}

	// looping a slice

	number := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range number {

		fmt.Println("index", index, "value", value)
	}
	// loopinf over a map

	//  looping over a string

	word := "Golang"
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	//conetinye and break

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skip the current iteration
		}
		if i == 5 {
			break // Exit the loop
		}
		fmt.Println(i)
	}

}
