package main

import (
	"fmt"
	"sync"
)

func fromSliceToChannel1(x []int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c)
	for _, val := range x {
		c <- val
	}
}

func fromChannel1ToChannel2(c1 <-chan int, c2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c2)

	for x := range c1 {
		c2 <- x * 2
	}
}

func fromChannel2ToOutput(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range c {
		fmt.Println(x)
	}
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := make(chan int)
	c2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)
	go fromSliceToChannel1(x, c1, &wg)
	go fromChannel1ToChannel2(c1, c2, &wg)
	go fromChannel2ToOutput(c2, &wg)
	wg.Wait()
}
