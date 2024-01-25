package main

import "fmt"

func main() {
	// print(5)
	// printRev(5)
	printBoth(5)
}

func print(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	print(n - 1)
}

func printRev(n int) {
	if n == 0 {
		return
	}
	printRev(n - 1)
	fmt.Println(n)

}

func printBoth(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printRev(n - 1)
	fmt.Println(n)

}
