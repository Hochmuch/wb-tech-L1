package main

import (
	"fmt"
	"sort"
)

func intersect(a []int, b []int) []int {
	result := make([]int, 0)
	sort.Ints(a)
	sort.Ints(b)

	n, m := len(a), len(b)
	i, j := 0, 0

	for i < n && j < m {
		if a[i] == b[j] {
			if len(result) == 0 || result[len(result)-1] != a[i] {
				result = append(result, a[i])
			}
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(intersect(a, b))
}
