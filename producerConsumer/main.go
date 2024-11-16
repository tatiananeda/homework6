package main

import (
	"fmt"
	"sync"
)

func producer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func consumer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		fmt.Println("Consumer processed: ", v)
	}
}

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	wg.Add(1)
	go producer(c, &wg)

	wg.Add(1)
	go consumer(c, &wg)

	wg.Wait()
}
