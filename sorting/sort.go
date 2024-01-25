package main

import "fmt"

func main() {
	arr := []int{4, 5, 1, 2, 3}
	// selection(arr)
	// insetion(arr)
	bubble(arr)
	fmt.Println(arr)
}

func reSelection(arr []int){
	
}

////////////////////////////////////////Selection Sort////////////////////////////////////
func selection(arr []int) {
	for i := 0; i < len(arr); i++ {
		last := len(arr) - i - 1
		maxIndex := getMaxIndex(arr, 0, last)
		arr[maxIndex], arr[last] = arr[last], arr[maxIndex]
	}
	return
}

func getMaxIndex(arr []int, start int, end int) int {
	max := start
	for i := start; i < end; i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}
	return max
}

// func swap(arr []int, first int, second int) {
// 	temp := arr[first]
// 	arr[first] = arr[second]
// 	arr[second] = temp
// }

//////////////////////////////////////Insetion Sort/////////////////////////////////////////////

func insetion(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
