package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

/*
Обоснование выбора способа завершения. Context то решение,
специально придуманное для этой задачи. И так как оно есть в стандартной библиотеке
лучше его и использовать. Это лучше, чем писать в канал, потому что context и есть
обёртка над каналом, где всё удобно реализовано.
*/

func recordConstantly4(ctx context.Context, channel chan int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				return
			}
			channel <- i
			i++
		}
	}
}

func worker4(ctx context.Context, id int, channel chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				return
			}
			x := <-channel
			fmt.Printf("Worker #%v прочитал: %v\n", id, x)
		}
	}
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	channel := make(chan int)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		fmt.Println("Прерывание Ctrl-C")
		cancel()
	}()

	for i := 0; i < n; i++ {
		go worker4(ctx, i, channel)
	}

	recordConstantly4(ctx, channel)

}
