package main

import "fmt"

func deleteElement(s []int, pos int) []int {
	copy(s[pos:], s[pos+1:])
	return s[:len(s)-1]
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = deleteElement(s, 2)
	fmt.Println(s)
}
