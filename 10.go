package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mp := make(map[int][]float64)

	for _, val := range data {
		intPart := int(val) / 10
		mp[intPart] = append(mp[intPart], val)
	}

	keys := make([]int, 0)
	for key := range mp {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, mp[key])
	}
}
