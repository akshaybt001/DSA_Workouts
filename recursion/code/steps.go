package main

import "fmt"

func numberOfSteps(num int) int {
	steps := 0
	return helper(num, steps)
}

func helper(num int, steps int) int {
	if num == 0 {
		return steps
	}
	if num%2 == 0 {
		return helper(num/2, steps+1)
	}
	return helper(num-1, steps+1)
}

func main() {
	num := numberOfSteps(41)
	fmt.Println(num)
}