package main

import (
	"fmt"
	"sync"
)

func errorOut(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered: ", err)
			wg.Done()
		}
	}()

	panic("Something went wrong")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("here comes the panic")

	go errorOut(&wg)

	wg.Wait()

	fmt.Println("recovered :)")
}
