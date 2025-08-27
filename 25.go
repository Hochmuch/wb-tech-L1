package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	<-timer.C
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sleep(time.Second * 5)
		fmt.Println("Прошло 5 секунд")
	}()
	for i := 1; i <= 4; i++ {
		sleep(time.Second)
		fmt.Printf("Прошло %v секунд\n", i)

	}
	wg.Wait()
}
