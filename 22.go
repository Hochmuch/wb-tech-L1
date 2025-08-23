package main

import (
	"fmt"
	"math/big"
)

func main() {
	var rawA, rawB = "5555555555555", "9999999999999"
	var a, b, c big.Int
	a.SetString(rawA, 10)
	b.SetString(rawB, 10)

	fmt.Println(c.Mul(&a, &b))
	fmt.Println(c.Div(&a, &b))
	fmt.Println(c.Add(&a, &b))
	fmt.Println(c.Sub(&a, &b))
}
