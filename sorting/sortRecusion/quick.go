package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func sort(num []int, low int, hi int) {
	if low >= hi {
		return
	}
	s := low
	e := hi
	m := s + (e-s)/2
	pivot := num[m]
	for s <= e {
		for num[s] < pivot {
			s++
		}
		for num[e] > pivot {
			e--
		}
		if s <= e {
			num[s], num[e] = num[e], num[s]
			s++
			e--
		}
	}
	sort(num, low, e)
	sort(num, s, hi)
}
