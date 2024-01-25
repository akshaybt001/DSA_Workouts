package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type doubleLinkedList struct {
	head *node
	tail *node
}

func (dl *doubleLinkedList) prepend(value int) {
	newNode := &node{data: value}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
		return
	}
	newNode.next = dl.head
	dl.head.prev = newNode
	dl.head = newNode

}

func (dl *doubleLinkedList) display() {
	current := dl.head
	for current != nil {
		fmt.Printf("%d==>", current.data)
		current = current.next
	}
}

func (dl *doubleLinkedList) revdisplay() {
	current := dl.tail
	for current != nil {
		fmt.Printf("%d==>", current.data)
		current = current.prev
	}
}

func main() {
	mylist := doubleLinkedList{}
	mylist.prepend(5)
	mylist.prepend(4)
	mylist.prepend(3)
	mylist.prepend(2)

	// mylist.display()
	mylist.revdisplay()

}
