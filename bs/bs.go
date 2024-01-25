package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 83}
	target := 9
	result := search(arr, target, 0, len(arr)-1)
	fmt.Println(result)
}

func search(arr []int, target int, s int, e int) int {
	if s <= e {
		mid := (s + e) / 2
		fmt.Println(mid, "k")
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
