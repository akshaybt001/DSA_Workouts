package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head *node
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

//////////////////////////////////////////////

func (l *linkedList) prepend(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

//////////////////////////////////////////////

func (l *linkedList) deleteTarget(target int) {
	if l.head == nil {
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == target {
			current.next = current.next.next

		}
		current = current.next

	}
}

////////////////////////////////////////////////

func (l *linkedList) deleteBefore(target int) {
	if l.head == nil {
		return
	}
	prev := l.head
	current := l.head
	for current.next != nil {
		if current.next.data == target {
			prev.next = prev.next.next
		}
		prev = current
		current = current.next
	}

}

///////////////////////////////////////////////
func (l *linkedList) deleteAfter(target int) {
	if l.head == nil {
		return
	}
	current := l.head
	for current.next != nil {
		if current.data == target {
			current.next = current.next.next
		}
		current = current.next
	}
}

func (l *linkedList) Print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d==>", current.data)
		current = current.next
	}
	fmt.Println("end")

}

///////////////////////////////////////////////

func (l *linkedList) insertAfter(value int, target int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
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

///////////////////////////////////////////////////
func (l *linkedList) insertBefore(value int, target int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
	}
	if l.head.data == target {
		newNode.next = l.head
		l.head = newNode
	}
	current := l.head
	for current != nil {
		if current.next.data == target {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}

////////////////////////////////////////////////////////
func (l *linkedList) findMid() int {
	if l.head == nil {
		return 0
	}
	fast := l.head
	slow := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.data
}

///////////////////////////////////////////////////
func (l *linkedList) arrtoLink(arr []int) {
	for _, val := range arr {
		l.append(val)
	}
}

func main() {
	mylist := linkedList{}
	mylist.append(5)
	mylist.append(4)
	mylist.append(3)
	mylist.append(2)
	// mylist.prepend(6)
	mylist.Print()
	// mylist.deleteTarget(3)
	// mylist.deleteBefore(3)
	// mylist.deleteAfter(5)
	// mylist.insertAfter(5, 2)
	// mylist.insertBefore(5, 3)
	hi := mylist.findMid()
	fmt.Println(hi)
	arr := []int{10, 20, 30, 40, 50}
	mylist.arrtoLink(arr)

	mylist.Print()
}
