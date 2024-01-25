package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

}

func (q *Queue) IsEmpty() bool {
	if q.front == nil {
		return true
	}
	return false
}

func (q *Queue) Print() {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		fmt.Printf("%d\t", current)
		temp.Enqueue(current)
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}

}
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	remove := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return remove
}

type Stack struct {
	q1, q2 *Queue
}

func (s *Stack) Push(value int) {
	for !s.q1.IsEmpty() {
		s.q2.Enqueue(s.q1.Dequeue())
	}
	s.q1.Enqueue(value)
	for !s.q2.IsEmpty() {
		s.q1.Enqueue(s.q2.Dequeue())
	}
}

func (s *Stack) Print() {
	s.q1.Print()
}

func NewStack() *Stack {
	return &Stack{q1: &Queue{}, q2: &Queue{}}
}

func main() {
	// q := Queue{}
	// q.Enqueue(10)
	// q.Enqueue(20)
	// q.Enqueue(30)
	// q.Enqueue(40)
	// q.Enqueue(50)
	// q.Print()
	s := NewStack()
	s.Push(50)
	s.Push(40)
	s.Print()
}
