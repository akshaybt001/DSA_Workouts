package main

import "fmt"

func count(num int) int {
	if num == 0 {
		return 1
	}
	if num < 10 {
		return 0
	}
	if num%10 == 0 {
		return 1 + count(num/10)
	} else {
		return count(num / 10)
	}
}

func main() {
	arr := 10250608
	zeroCount := count(arr)
	fmt.Println("Number of zero in ", zeroCount)
}
