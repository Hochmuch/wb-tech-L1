package main

import "fmt"

func binarySort(a []int, x int) int {
	l, r := 0, len(a)
	for l+1 < r {
		m := (l + r) / 2

		if a[m] > x {
			r = m
		} else {
			l = m
		}
	}
	if a[l] == x {
		return l
	}
	return -1
}

func main() {
	fmt.Println(binarySort([]int{1, 2, 3, 4, 5, 6}, 4))
}
