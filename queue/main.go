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
	newNode := &Node{data: value, next: nil}
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

func (q *Queue) Dequeus() int {
	if q.IsEmpty() {
		return -1
	}
	toremove := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return toremove
}

func (q *Queue) Print() {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeus()
		fmt.Printf("%d\t", current)
		temp.Enqueue(current)
	}
	fmt.Printf("\n")
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeus())
	}
}

func (q *Queue) Size() int {
	size := 0
	temp := Queue{}
	for !q.IsEmpty() {
		temp.Enqueue(q.Dequeus())
		size++
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeus())
	}
	return size
}

func (q *Queue)insertAfter(value,target int){
	temp:=Queue{}
	for !q.IsEmpty(){
		current:=q.Dequeus()
		temp.Enqueue(current)
		if current==target{
			temp.Enqueue(value)
		}
	}
	for !temp.IsEmpty(){
		q.Enqueue(temp.Dequeus())
	}
}


func main() {
	q := Queue{}
	q.Enqueue(50)
	q.Enqueue(60)
	q.Enqueue(70)
	q.Enqueue(80)
	q.Print()
	fmt.Println(q.Size())

}
