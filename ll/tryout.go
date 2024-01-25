package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) append1(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *linkedList) append(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *linkedList) Print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d =>", current.data)
		current = current.next
	}
	fmt.Printf("End")

}

func (l *linkedList) arrtoLink(arr []int) {
	for _, val := range arr {
		l.append(val)
	}
}

func (l *linkedList) insertAfter(target int, value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current != nil {
		if current.data == target {
			newNode.next = current.next
			current.next = newNode
		}
		current = current.next
	}
}

func (l *linkedList) deleteBefore(target int) {
	if l.head == nil {
		return
	}
	if l.head.next.data == target {
		s:=l.head.next
		l.head = s
		return
	}
	current := l.head
	prev := l.head
	for current != nil {
		if current.next.data == target {
			prev.next = current.next
			break
		}
		prev = current

		current = current.next
	}
}

func main() {
	mylist := linkedList{}
	arr := []int{1, 15, 16, 25, 99, 66}
	// mylist.append(1)
	mylist.arrtoLink(arr)
	// mylist.insertAfter(25, 55)
	mylist.deleteBefore(15)
	mylist.Print()

}
