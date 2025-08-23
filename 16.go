package main

import "fmt"

func quickSort(a []int) []int {
	n := len(a)

	if n < 2 {
		return a
	}

	pivot := 0

	for i := 1; i < n; i++ {
		if i == pivot {
			continue
		}

		if i < pivot && a[i] > a[pivot] {
			a[i], a[pivot] = a[pivot], a[i]
		}

		if i > pivot && a[i] < a[pivot] {
			a[i], a[pivot] = a[pivot], a[i]
		}
	}
	quickSort(a[:pivot])
	quickSort(a[pivot+1:])

	return a
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, -2025}
	a = quickSort(a)
	fmt.Println(a)
}
