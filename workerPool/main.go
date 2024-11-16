package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, queue chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range queue {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d processing job %d \n", id, val)
	}
}

func dispatcher(jobs int, workers int) {
	queue := make(chan int, jobs)

	var wg sync.WaitGroup

	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go worker(i, queue, &wg)
	}

	for i := 1; i <= jobs; i++ {
		queue <- i
	}

	close(queue)

	wg.Wait()
}

func main() {
	jobs := 5
	workers := 3

	// submit jobs to workers
	dispatcher(jobs, workers)
}
