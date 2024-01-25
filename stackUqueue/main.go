package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{data: value, next: s.top}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Pop() int {
	pop := s.top.data
	s.top = s.top.next
	return pop
}

func (s *Stack) Print() {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		fmt.Printf("%d\n", current)
		temp.Push(current)
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}

}

type Queue struct {
	s1, s2 *Stack
}

func NewQueue() *Queue {
	return &Queue{s1: &Stack{}, s2: &Stack{}}
}

func (q *Queue) Enqueue(value int) {
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
	q.s1.Push(value)
	for !q.s2.IsEmpty() {
		q.s1.Push(q.s2.Pop())
	}
}

func (q *Queue) Dequeue() {
	q.s1.Pop()
}
func (q *Queue) Print() {
	q.s1.Print()
}

func main() {
	// myStack := Stack{}
	// myStack.Push(30)
	// myStack.Push(20)
	// myStack.Push(10)
	// myStack.Push(5)
	// myStack.Print()
	q := NewQueue()
	q.Enqueue(50)
	q.Enqueue(60)
	q.Enqueue(70)
	q.Print()
}
