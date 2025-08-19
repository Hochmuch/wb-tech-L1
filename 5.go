package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func write(c chan int) {
	i := 0
	for {
		c <- i
	}
}

func read(c chan int) {
	var x int
	for {
		x = <-c
		fmt.Println("Считано:", x)
	}
}

func main() {
	N, _ := strconv.Atoi(os.Args[1])

	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)

	go write(c)
	go read(c)

	<-time.After(time.Duration(N) * time.Second)

	wg.Done()
	wg.Done()
	fmt.Printf("Прошло %v секунд\n", N)
}
