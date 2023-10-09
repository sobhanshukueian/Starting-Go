package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	// Create nodes and link them
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	node1.Next = node2
	node2.Next = node3

	// Traverse the linked list
	currentNode := node1
	for currentNode != nil {
		fmt.Println("Node Value:", currentNode.Value)
		currentNode = currentNode.Next
	}
}
