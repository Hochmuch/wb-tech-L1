package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func byCondition(flag bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if flag {
		fmt.Println("Выход из горутины по условию")
		return
	}
}

func byNotificationChannel(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-c:
			fmt.Println("Выход из горутины через канал уведомления")
			return
		default:

		}
	}
}

func byContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Выход из горутины через контекст")
			return
		default:

		}
	}
}

func byRuntimeGoexit(flag bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if flag {
			fmt.Println("Выход с помощью runtime.Goexit()")
			runtime.Goexit()
		}
	}
}

func main() {
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(4)

	go byCondition(true, &wg)
	go byNotificationChannel(c, &wg)
	close(c)
	go byContext(ctx, &wg)
	cancel()
	go byRuntimeGoexit(true, &wg)

	wg.Wait()

}
