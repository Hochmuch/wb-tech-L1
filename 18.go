package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	x  int
	mu sync.Mutex
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	for i := 0; i < 1000; i++ {
		c.x++
	}
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var c = &Counter{x: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.increment(&wg)
	}

	wg.Wait()
	fmt.Println(c.x)
}
