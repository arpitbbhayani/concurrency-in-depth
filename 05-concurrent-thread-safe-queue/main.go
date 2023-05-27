package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	lock  sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.queue) == 0 {
		panic("removing from an empty queue")
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue)
}

const NUM_THREADS int = 1000000

func main() {
	queue := &ConcurrentQueue{}

	var wg sync.WaitGroup

	for i := 0; i < NUM_THREADS; i++ {
		go func() {
			wg.Add(1)
			queue.Enqueue(rand.Int31())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("size:", queue.Size())

	for i := 0; i < NUM_THREADS; i++ {
		go func() {
			wg.Add(1)
			queue.Dequeue()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("size:", queue.Size())
}
