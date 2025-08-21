package main

import (
	"fmt"
	"sort"
)

func makeSet(s []string) []string {
	result := make([]string, 0)
	sort.Strings(s)

	for _, val := range s {
		if len(result) == 0 || val != result[len(result)-1] {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	fmt.Println(makeSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}
