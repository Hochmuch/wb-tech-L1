package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
