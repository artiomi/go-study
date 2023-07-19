package main

import (
	"fmt"
	"time"
)

func main() {

	var myChannel = make(chan string)
	var anotherChannel = make(chan string)
	go func() {
		myChannel <- "cat"
		fmt.Println("sending data to myChannel")
	}()

	go func() {
		anotherChannel <- "dog"
		fmt.Println("sending data to anotherChannel")
	}()
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Printf("Msg from myChannel: %v", msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Printf("Msg from anotherChannel: %v", msgFromAnotherChannel)
	}
	time.Sleep(10 * time.Second)
}
