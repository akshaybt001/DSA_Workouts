package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{value, s.top}
	s.top = newNode

}

func (s *Stack) Print() {
	current := s.top
	fmt.Print("Stack: ")
	for current != nil {
		fmt.Printf("%d \n", current.data)
		current = current.next
	}
}

func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return 0
	}
	popped := s.top.data
	s.top = s.top.next
	return popped

}

func (s *Stack) Peek() int {
	if s.top == nil {
		return 0
	}
	return s.top.data
}

func (s *Stack) size() int {
	if s.top == nil {
		return 0
	}
	size := 0
	current := s.top
	for current != nil {
		size++
		current = current.next
	}
	return size
}
func (s *Stack) clean() {
	s.top = nil
}

func (s *Stack) contain(t int) bool {
	current := s.top
	for current != nil {
		if current.data == t {
			return true
		}
		current = current.next
	}
	return false
}

func (s *Stack) DeleteWithValue(value int) {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		if current != value {
			temp.Push(current)
		}
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}

func (s *Stack) InsertBefore(value int, target int) {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		temp.Push(current)
		if current == target {
			temp.Push(value)
		}
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}

func (s *Stack) InsertAfter(value, target int) {
	temp := Stack{}
	for !s.IsEmpty() {
		if s.Peek() == target {
			temp.Push(value)
			break
		}
		current := s.Pop()
		temp.Push(current)
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	myStack.Print()
	// myStack.Pop()
	// myStack.Print()

	//
	// peek1 := myStack.Peek()
	// fmt.Println("peeked value: ", peek1)
	// fmt.Println("size:", myStack.size())

	// fmt.Println(":", myStack.contain(50))
	// myStack.DeleteWithValue(20)
	// myStack.InsertBefore(88, 20)
	myStack.InsertAfter(88, 20)
	myStack.Print()
	// myStack.clean()

}
