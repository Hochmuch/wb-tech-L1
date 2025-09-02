package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	s = strings.ToLower(s)
	mp := make(map[rune]int)
	for _, c := range []rune(s) {
		_, ok := mp[c]
		if ok {
			return false
		}
		mp[c]++
	}
	return true
}

func main() {
	flag := check("abcdC")
	if flag {
		fmt.Println("All letters are unique")
	} else {
		fmt.Println("Not all letters are unique")
	}
}
