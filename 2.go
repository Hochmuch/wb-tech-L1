package main

import (
	"fmt"
	"sync"
)

func pow2(x int) {
	fmt.Println(x * x)
}

func main() {
	var wg sync.WaitGroup

	nums := [...]int{2, 4, 6, 8, 10}
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			pow2(num)
			defer wg.Done()
		}(num)

	}
	wg.Wait()
}
