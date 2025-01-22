package main

import (
	"fmt"
)

type Node struct {
	value int
	Next  *Node
}

func main() {

	node1 := &Node{value: 1}

	node2 := &Node{value: 2}

	node3 := &Node{value: 4}

	node1.Next = node2

	node2.Next = node3

	current := node1

	for current != nil {
		fmt.Println(current.value)
		current = current.Next

	}
}
