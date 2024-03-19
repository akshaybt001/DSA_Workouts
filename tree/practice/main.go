package main

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

func (t *Tree) Delete(value int){
	t.root= delete(t.root,value)
}

func delete(root *Node,value int)*Node{
	if root.key<value{
		root.right= delete(root.right,value)
	}else if root.key>value{
		root.left=delete(root.left,value)
	}else{
		if root.left==nil{
			return root.right
		}else if root.right==nil{
			return root.left
		
		}
		root.key=inorderSuccessor(root.right)
		root.right=delete(root.right,root.key)
	}
	return root
}

func inorderSuccessor(root *Node)int{
	current:=root
	for current.left!=nil{
		current=current.left
	}
	return current.key
}