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
	queue := &ConcurrentQueue{
		queue: make([]int32, 0),
	}

	var wgE sync.WaitGroup
	var wgD sync.WaitGroup

	for i := 0; i < NUM_THREADS; i++ {
		wgE.Add(1)
		go func() {
			queue.Enqueue(rand.Int31())
			wgE.Done()
		}()
	}

	for i := 0; i < NUM_THREADS; i++ {
		wgD.Add(1)
		go func() {
			queue.Dequeue()
			wgD.Done()
		}()
	}

	wgE.Wait()
	wgD.Wait()

	fmt.Println("size:", queue.Size())
}
