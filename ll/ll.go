package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head *node
}

//====================== append =========================
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

//===================== Prepend ======================
func (l *linkedList) prepend(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

// ====================== Print ============================
func (l *linkedList) Print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ==>", current.data)
		current = current.next
	}
}

// =======================Delete First =======================
func (l linkedList) deleteFirst() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

//====================== Delete last ========================

func (l linkedList) deleteLast() {
	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

//==========================Insert After =====================

func (l linkedList) insertAfter(T int, value int) {
	newNode := &node{data: value}
	current := l.head
	for current != nil {
		if current.data == T {
			newNode.next = current.next
			current.next = newNode
		}
		current = current.next
	}
}

// ======================Insert Before =========================
func (l linkedList) insertBefore(T int, value int) {
	newNode := &node{data: value}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current.data == T {
			newNode.next = current
			prev.next = newNode
		}
		prev = current
		current = current.next
	}
}

//=========================Delete Target =====================
func (l linkedList) deleteTarget(T int) {
	current := l.head
	for current.next != nil {
		if current.next.data == T {
			current.next = current.next.next
		}
		current = current.next
	}
}

func main() {
	l := linkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	// l.prepend(5)
	// l.deleteFirst()
	// l.Print()
	// l.deleteLast()
	// l.insertAfter(2, 4)
	l.deleteTarget(2)
	l.Print()
}
