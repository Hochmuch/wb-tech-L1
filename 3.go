package main

import (
	"fmt"
	"os"
	"strconv"
)

func recordConstantly(channel chan int) {
	i := 0
	for {
		channel <- i
		i++
	}
}

func worker(id int, channel chan int) {
	for {
		x := <-channel
		fmt.Printf("Worker #%v прочитал: %v\n", id, x)
	}
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	channel := make(chan int)

	for i := 0; i < n; i++ {
		go worker(i, channel)
	}

	recordConstantly(channel)

}
