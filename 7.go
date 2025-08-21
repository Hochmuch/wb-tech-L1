package main

import (
	"fmt"
	"sync"
)

/*
Здесь демонстрируется разница между конкурентой записью
в map с использованием mutex и без него.
Насколько я понял, смысл задания был в этом.
*/

type MapWithMutex struct {
	mp map[int]int
	mu sync.Mutex
}

func (m *MapWithMutex) incrementCorrectly() {
	m.mu.Lock()
	m.mp[0]++
	m.mu.Unlock()
}

func (m *MapWithMutex) incrementIncorrectly() {
	m.mp[1]++
}

func main() {
	var mwm MapWithMutex = MapWithMutex{mp: make(map[int]int)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mwm.incrementCorrectly()
				//mwm.incrementIncorrectly()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Ожидается 100000")
	fmt.Println("Результат при записи с mutex:", mwm.mp[0])
	// Если пытаться запустить неправильную запись, просто получаю fatal error: concurrent map writes
}
