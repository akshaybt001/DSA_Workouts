package main

import "fmt"

func search(arr []int, target int, s int, e int) int {
	if s <= e {
		mid := (e + s) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return search(arr, target, mid+1, e)
		} else {
			return search(arr, target, s, mid-1)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 7
	result := search(arr, target, 0, len(arr)-1)
	fmt.Println(result)
}
