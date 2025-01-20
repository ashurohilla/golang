package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello i am there and lets code")

	maper := make(map[string]int)

	maper["ashish"] = 25
	maper["raghav"] = 36

	fmt.Println(maper)

	// another method to impements maps in go lang

	mapernew := map[string]string{
		"alice": "hello",
		"data":  "bob",
	}

	fmt.Println(mapernew)

	//checking value of a map key

	fmt.Println(mapernew["alice"])

	//checking for key of a map

	age, exist := mapernew["random"]

	if exist {
		fmt.Println(age)
	} else {
		fmt.Println("value not found in the mapermew")
	}

	// iterating over  a maps and getting data .

	for key, value := range mapernew {

		fmt.Printf("key: %s , value : %s\n", key, value)
	}
}
