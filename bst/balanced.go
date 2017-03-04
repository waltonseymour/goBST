package bst

import (
	"errors"
	"sort"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// assumes sorted input
func treeHelper(nums ...int) *BST {
	if !sort.IntsAreSorted(nums) {
		errors.New("Input should be sorted")
	}
	if len(nums) == 0 {
		return &BST{nil, 0, 0}
	}

	index := len(nums) / 2
	newNode := Node{val: nums[index]}
	bst := BST{root: &newNode, size: 1, depth: 1}
	if index != 0 {
		leftSubTree := treeHelper(nums[:index]...)
		rightSubTree := treeHelper(nums[index+1:]...)
		newNode.left = leftSubTree.root
		newNode.right = rightSubTree.root
		bst.size += leftSubTree.size + rightSubTree.size
		bst.depth += max(leftSubTree.depth, rightSubTree.depth)
	}
	return &bst
}

// BalancedTree returns a new Balanced BST for a given input
func BalancedTree(nums ...int) *BST {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	nums = sorted
	return treeHelper(nums...)
}
