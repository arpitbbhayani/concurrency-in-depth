package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentQueue struct {
	queue []int32
	lock  sync.Mutex
	cv    *sync.Cond
}

func NewConcurrentQueue() *ConcurrentQueue {
	tsqueue := ConcurrentQueue{
		queue: make([]int32, 0),
		lock:  sync.Mutex{},
	}
	tsqueue.cv = sync.NewCond(&tsqueue.lock)
	return &tsqueue
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, item)
	q.cv.Signal() // This can be doen outside lock
}

func (q *ConcurrentQueue) Dequeue() int32 {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.queue) == 0 {
		fmt.Println("queue is empty so consumer will wait until notified")
		q.cv.Wait() // Unlocks and waits, when notified locks and wakesup
		// After wakeup check again if queue is not empty
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

const NUM_PRODUCERS int = 10
const NUM_CONSUMERS int = 10

func main() {
	queue := NewConcurrentQueue()

	var wgE sync.WaitGroup
	var wgD sync.WaitGroup

	for i := 0; i < NUM_PRODUCERS; i++ {
		wgE.Add(1)
		go func() {
			for {
				queue.Enqueue(rand.Int31())
				time.Sleep(time.Second * 10)
				// Producers will produce data every 10 secs
				// Consumers will wait when all the data on queue is consumed
			}
			wgE.Done()
		}()
	}

	for i := 0; i < NUM_CONSUMERS; i++ {
		wgD.Add(1)
		go func() {
			for {
				item := queue.Dequeue()
				fmt.Println("Consumed: ", item)
			}
			wgD.Done()
		}()
	}

	wgE.Wait()
	wgD.Wait()

	fmt.Println("size:", queue.Size())
}
