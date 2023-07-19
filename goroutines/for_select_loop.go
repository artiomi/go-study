package main

import "fmt"

func main() {
	buffChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	for _, char := range chars {
		select {
		case buffChannel <- char:
		}
	}
	close(buffChannel)
	for result := range buffChannel {
		fmt.Printf("Channel value: %v.\n", result)
	}
}
