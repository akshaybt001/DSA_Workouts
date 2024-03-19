package main

import (
	"fmt"
)

func reverseString(s []byte) {
	helper(s, 0, len(s)-1)
}

func helper(s []byte, L int, R int) {
	if L >= R {
		return
	}
	s[L], s[R] = s[R], s[L]
	helper(s, L+1, R-1)
}

func main() {
	str := "akshay"
	bytey := []byte(str)
	reverseString(bytey)
	fmt.Println(string(bytey))
}
