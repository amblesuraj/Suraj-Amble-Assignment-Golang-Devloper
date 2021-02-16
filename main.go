package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n************ Assignment 1 - Binary Tree Assignment *******************\n")

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{60, nil, nil}
	root.right.left = &Node{50, nil, nil}

	fmt.Println("\nIn Order Traversal")
	InOrderTraversal(root)

	fmt.Println("\nPreorder Traversal")
	PreOrderTraversal(root)

	fmt.Println("\nPost Order Traversal")
	PostOrderTraversal(root)

	fmt.Println("\n\n***************************************************\n")
	fmt.Println("\n***************** Assignment 2 - House Robber ********************")

	highValues := []float64{2, 3, 2}
	highValues1 := []float64{1, 3, 2, 1}
	highValues2 := []float64{0}

	fmt.Println("\nHigh Values sum from {2,3,2} ", Rob(highValues...))
	fmt.Println("\nHigh Values sum from {1,3,2,1} ", Rob(highValues1...))
	fmt.Println("\nHigh Values sum from {0} ", Rob(highValues2...))

	fmt.Println("\n***************************************************\n")
}
