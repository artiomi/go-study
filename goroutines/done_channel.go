package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go doWork(done)
	time.Sleep(5 * time.Second)
	println("closing channel")
	close(done)

}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work.")
		}
	}
}
