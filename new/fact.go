package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}

}



func main() {
	result := fact(5)
	fmt.Println(result)
}
