package main

import "fmt"

func reverseString(str []byte, sr int, e int) {
	if sr >= e {
		return
	}
	str[sr], str[e] = str[e], str[sr]
	reverseString(str, sr+1, e-1)
}

func helper(s string) string {
	str := []byte(s)
	reverseString(str, 0, len(s)-1)
	return string(str)
}

func main() {
	str := "akshay"
	hai := helper(str)
	fmt.Println(hai)
}
