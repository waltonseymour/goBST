package bst

// BinaryTree returns a new non-balanced BST
func BinaryTree(nums ...int) *BST {
	root := Node{val: nums[0]}
	bst := BST{root: &root, size: 1, depth: 1}
	for i, num := range nums {
		if i == 0 {
			continue
		}

		// traverses bst for given input
		node, depth := bst.traverse(num)

		// constructs a new node
		newNode := Node{val: num}

		// sets as right child if greater
		// and left if less than
		// ignores duplicte entries
		if num > node.val {
			node.right = &newNode
		} else if num < node.val {
			node.left = &newNode
		} else {
			continue
		}

		// if we've exceeded our max depth, we increase
		// the max depth of the BST
		if depth > bst.depth {
			bst.depth = depth
		}

		// we increase our BST size
		bst.size++

	}
	return &bst
}
