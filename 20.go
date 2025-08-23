package main

import "fmt"

func reverseSentence(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	l := 0
	for i := 0; i <= n; i++ {
		if i == n || r[i] == ' ' {
			for j := 0; j < (i-l)/2; j++ {
				r[l+j], r[i-1-j] = r[i-1-j], r[l+j]
			}
			l = i + 1
		}
	}
	return string(r)
}

func main() {
	fmt.Println(reverseSentence("snow dog sun"))
}
