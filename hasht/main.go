package main

import (
	"fmt"
)

const ArraySize = 9

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].Search(key)

}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// Insert will take in a key , create a node with the key and insert th enode in the bucket
func (b *bucket) insert(k string) {
	if !b.Search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// Search wil take in a  key and return true if the buket has that key
func (b *bucket) Search(k string) bool {

	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

// Delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prev := b.head
	for prev.next != nil {
		if prev.next.key == k {
			//delete
			prev.next = prev.next.next
		}
		prev = prev.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	hashTable.Insert("RANDY")
	hashTable.Insert("KANE")
	// hashTable.Delete("RANDY")
	fmt.Println(hashTable.Search("RANDY"))
	fmt.Println(hashTable.Search("Akshay"))

}
