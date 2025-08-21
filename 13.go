package main

import "fmt"

func swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

func main() {
	a, b := 1337, 2025
	swap(&a, &b)
	fmt.Println(a, b)
}
