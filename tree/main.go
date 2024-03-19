package main

import (
	"fmt"
	"math"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	t.root = insert(t.root, value)
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{key: value}
	}
	if root.key < value {
		root.right = insert(root.right, value)
	} else {
		root.left = insert(root.left, value)
	}
	return root
}

func (t *Tree) Search(value int) bool {
	return search(t.root, value)
}

func search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.key < value {
		return search(root.right, value)
	} else if root.key > value {
		return search(root.left, value)
	}
	return true
}

func (t *Tree) Delete(value int) {
	t.root = delete(t.root, value)
}

func delete(root *Node, value int) *Node {
	if root.key > value {
		root.left = delete(root.left, value)
	} else if root.key < value {
		root.right = delete(root.right, value)
	} else {
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		root.key = inorderSuccessor(root.right)
		root.right = delete(root.right, root.key)
	}
	return root
}

func inorderSuccessor(root *Node) int {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current.key
}
func InorderTranversal(root *Node) {
	if root != nil {
		InorderTranversal(root.left)
		fmt.Printf("%d => ", root.key)
		InorderTranversal(root.right)
	}
}

func PreOrderTranversal(root *Node) {
	if root != nil {
		fmt.Printf("%d =>", root.key)
		PreOrderTranversal(root.left)
		PreOrderTranversal(root.right)
	}
}

func PostOrderTraversal(root *Node) {
	if root != nil {
		PostOrderTraversal(root.left)
		PostOrderTraversal(root.right)
		fmt.Printf("%d =>", root.key)
	}
}
func IsValidBst(node *Node) bool {
	return isValidBst(node, math.MinInt64, math.MaxInt64)
}

func isValidBst(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.key <= min || node.key >= max {
		return false
	}
	return isValidBst(node.left, min, node.key) && isValidBst(node.right, node.key, max)
}

func main() {
	t := Tree{}
	arr := []int{100, 18, 22, 23, 44, 16, 101}
	for _, num := range arr {
		t.Insert(num)
	}
	fmt.Println(inorderSuccessor(t.root))
	PostOrderTraversal(t.root)
	fmt.Println(t.Search(4))
}
