package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Working on ", n)
	time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Started")
	defer fmt.Println("Completed")

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()
}
