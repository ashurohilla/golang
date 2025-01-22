package main

import (
	"fmt"
)

func main() {

	x := 24
	p := &x

	fmt.Println(p)

	*p = 40
	fmt.Println(x)

}
