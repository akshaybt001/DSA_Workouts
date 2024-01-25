package main

import "fmt"

func main() {
	ans := sum(1342)
	fmt.Println(ans)
}

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sum(n/10)
}
