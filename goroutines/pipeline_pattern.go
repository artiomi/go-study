package main

import "fmt"

func main() {
	//input
	nums := []int{2, 3, 4, 7, 1}
	//stage 1
	dataChannel := sliceToChannel(nums)
	//stage 2
	finalChannel := sq(dataChannel)
	//stage 3
	for n := range finalChannel {
		fmt.Printf("final value: %v\n", n)
	}
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for number := range in {
			fmt.Printf("Multipling value %v .\n", number)
			out <- number * number
		}
		println("Out channel closed in sq")
		close(out)
	}()

	return out
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range nums {
			fmt.Printf("Pushing value %v to the channel.\n", i)
			out <- i
		}
		println("out channel closed in sliceToChannel.")
		close(out)
	}()
	return out
}
