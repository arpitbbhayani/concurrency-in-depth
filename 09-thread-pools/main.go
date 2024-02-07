package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workQueue: make(chan Job),
	}
	pool.wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			for job := range pool.workQueue {
				job()
			}
		}()
	}
	return pool
}

func (p *Pool) AddJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5)

	for i := 0; i < 30; i++ {
		job := func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("job: completed\n")
		}
		pool.AddJob(job)
	}

	pool.Wait()
}
