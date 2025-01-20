package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("hello i am there and lets code")

	str := "ashish"

	str2 := "rohilla"

	result := str + str2
	fmt.Println(result)

	// lenght of a string

	length := len(result)
	fmt.Println(length)

	//accesing individual character

	fmt.Println(string(str[2]))

	//check for substrings

	hello := strings.Contains(result, "rohillaaa")

	fmt.Println(hello)

	//spliting strings
	parts := strings.Split(result, "a")
	fmt.Println(parts)

	// string conversions
	str4 := "123"

	newstr, err := strconv.Atoi(str4)
	if err != nil {

		fmt.Println("Error", err)
	} else {
		fmt.Println(newstr)

		fmt.Printf("type of newstr  %T\n", newstr)
	}
}
