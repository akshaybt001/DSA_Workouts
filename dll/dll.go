package main

import (
	"fmt"
)

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

func (dl *doubleLinkedList) append(value int) {
	newNode := &node{data: value}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
		return
	}
	current := dl.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = dl.tail
	dl.tail = newNode

}

func (dl doubleLinkedList) display() {
	current := dl.head
	for current != nil {
		fmt.Printf("%d =>", current.data)
		current = current.next
	}
	fmt.Println("End")
}
func (dl *doubleLinkedList) targetdelete(target int) {
	if dl.head == nil {
		return
	}
	if dl.head.data == target {
		dl.head = dl.head.next
		dl.head.prev = nil
		return
	}
	if dl.tail.data == target {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
		return

	}
	current := dl.head
	for current != nil {
		if current.next.data == target {
			current.next = current.next.next
			current.next.prev = current
			return

		}
		current = current.next
	}
}

func main() {
	mylist := doubleLinkedList{}
	mylist.append(8)
	mylist.append(2)
	mylist.append(1)
	mylist.append(3)

	mylist.append(5)
	mylist.append(6)
	mylist.append(9)
	mylist.display()
	mylist.targetdelete(5)
	mylist.display()

}
