package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int
var NUM_THREADS int = 1000000

func incCountWithoutLock(wg *sync.WaitGroup) {
	count++
	wg.Done()
}

func incCountWithMutex(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func countWithoutLock() {
	for i := 0; i < NUM_THREADS; i++ {
		wg.Add(1)
		go incCountWithoutLock(&wg)
	}
	wg.Wait()
}

func countWithMutex() {
	wg.Add(NUM_THREADS)
	for i := 0; i < NUM_THREADS; i++ {
		go incCountWithMutex(&wg)
	}
	wg.Wait()
}

func main() {
	count = 0
	countWithoutLock()
	fmt.Println("count (without lock)", count)

	count = 0
	countWithMutex()
	fmt.Println("count (with mutex)", count)
}
