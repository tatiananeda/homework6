package main

import "fmt"

func printNum(n int, done chan bool) {
	fmt.Println(n)
	done <- true
}

func main() {
	done := make(chan bool)
	defer close(done)

	fmt.Println("start")

	for i := 1; i <= 10; i++ {
		go printNum(i, done)
		<-done
	}

	fmt.Println("end")
}
