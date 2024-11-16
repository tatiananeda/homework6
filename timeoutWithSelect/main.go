package main

import (
	"fmt"
	"time"
)

func SendMsgInSec(sec int, msg string, c chan string) {
	defer close(c)
	time.Sleep(time.Duration(sec) * time.Second)
	c <- msg
}

func main() {
	c := make(chan string)

	go SendMsgInSec(2, "Waited two second", c)

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
