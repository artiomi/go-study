package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//doWalkSeq(t, ch)

	doWalkRecur(t, ch)

	close(ch)

}
func doWalkRecur(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		doWalkRecur(t.Left, ch)
	}
	if t.Right != nil {
		doWalkRecur(t.Right, ch)
	}
}

func doWalkSeq(t *tree.Tree, out chan int) {
	var counter uint = 0
	in := make(chan *tree.Tree, 1)
	in <- t
	counter++

	for {
		fmt.Println("counter is:", counter)
		if counter == 0 {
			fmt.Println("the end")
			close(in)
			break
		}

		t2 := <-in
		out <- t2.Value
		counter--
		if t2.Right != nil {
			go push(t2.Right, in)
			counter++

		}
		if t2.Left != nil {
			go push(t2.Left, in)
			counter++

		}
	}

}
func push(t *tree.Tree, ch chan *tree.Tree) {
	ch <- t
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	//fmt.Printf("Comparing t1: %v,\n t2: %v\n", t1, t2)
	if t1 == nil && t2 == nil {
		return true
	}
	//fmt.Printf("Values t1: %v, t2:%v \n", t1.Value, t2.Value)

	if t1.Value == t2.Value {
		return Same(t1.Right, t2.Right) && Same(t1.Left, t2.Left)
	}
	return false
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(2), ch)
	//for v := range ch {
	//	fmt.Println("received value:", v)
	//}
	t1 := tree.New(1)
	same := Same(t1, t1)
	fmt.Println("Are the same:", same)
}
