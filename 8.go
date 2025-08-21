package main

import "fmt"

func changeBit(num int64, pos int, to int64) int64 {
	bit := (num >> (pos - 1)) & 1
	if bit == to {
		return num
	}
	return num ^ (1 << (pos - 1))
}

func main() {
	fmt.Println(changeBit(8, 2, 1))
	fmt.Println(changeBit(8, 2, 0))
}
