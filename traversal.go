package main

import (
	"fmt"
)

// Node
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Preorder
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

// In Order
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%d ", root.val)
	InOrderTraversal(root.right)
}

// Post Order
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%d ", root.val)
}
