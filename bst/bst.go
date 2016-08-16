package bst

import (
  "math"
)

type Node struct {
  val int
  left, right *Node
}

type BST struct {
  root *Node
  size, depth int
}

// returns pointer to last node in traversal, looking for val
func (bst *BST) traverse(val int) (node *Node, depth int) {
  cur := bst.root
  depth = 0
  for {
    depth++
    if val > cur.val && cur.right != nil {
      cur = cur.right
    } else if val < cur.val && cur.left != nil {
      cur = cur.left
    } else {
      return cur, depth
    }
  }
}

func (bst *BST) Search(val int) bool {
  node, _ := bst.traverse(val)
  return node.val == val
}

func (bst *BST) RangeSearch(min int, max int) int {
  node, _ := bst.traverse(min)
  return node.val
}


func (bst *BST) Max() int {
  node, _ := bst.traverse(math.MaxInt64)
  return node.val
}

func (bst *BST) Min() int {
  node, _ := bst.traverse(0)
  return node.val
}
