package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"./bst"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	vals := []int{}

	for i := 0; i < 1000000; i++ {
		x := rand.Int()
		vals = append(vals, x)
	}

	var tree *bst.BST
	if len(os.Args) > 1 && os.Args[1] == "unbalanced" {
		tree = bst.BinaryTree(vals...)
	} else {
		tree = bst.BalancedTree(vals...)
	}
	fmt.Println("Created a BST with 1M random values")
	fmt.Println("Max: ", tree.Max())
	fmt.Println("Min: ", tree.Min())
	fmt.Println("Depth: ", tree.Depth())
}
