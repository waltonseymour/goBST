package bst

// non balancing
func BinaryTree(nums ...int) *BST {
  root := Node{val: nums[0]}
  bst := BST{root: &root, size: 1, depth: 1}
  for i, num := range nums {
    if i == 0 {
      continue
    } else {
      node, depth := bst.traverse(num)
      newNode := Node{val: num}
      if num > node.val {
        node.right = &newNode
      } else if num < node.val {
        node.left = &newNode
      } else {  // num == node.val
        continue
      }
      if depth > bst.depth{
        bst.depth = depth
      }
      bst.size++
    }
  }
  return &bst
}
