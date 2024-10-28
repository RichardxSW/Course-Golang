package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal completion to main goroutine
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(2 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal completion to main goroutine
	close(doneChan) // close the channel to prevent potential deadlock
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {

	}
}