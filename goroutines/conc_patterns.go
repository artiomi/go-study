package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	updateTicketStore()
}

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}

func Turnout(InA, InB <-chan int, OutA, OutB chan int) {
	var more bool
	var data int
	for {
		select { //Receive from first non-blocking
		case data, more = <-InA:
		case data, more = <-InB:
		}
		if !more {
			return
		}
		select { //Send to first non-blocking
		case OutA <- data:
		case OutB <- data:
		}
	}
}

type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string //for simplicity: imagine this to be infinite
}

func (ts *TicketStore) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1
	ts.slots[t] = s
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) {
		runtime.Gosched()
	}
}
func (ts *TicketStore) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1]

}

func updateTicketStore() {
	var ticket uint64 = 0
	var done uint64 = 0
	slots := []string{}
	ts := TicketStore{ticket: &ticket, done: &done, slots: slots}
	fmt.Printf("Initial store. ticket: %v done: %v slots: %v\n", *ts.ticket, *ts.done, ts.slots)
	ts.Put("v")
	fmt.Printf("after update store. ticket: %v done: %v slots: %v\n", *ts.ticket, *ts.done, ts.slots)
	ts.Put("k")
	fmt.Printf("after update store. ticket: %v done: %v slots: %v\n", *ts.ticket, *ts.done, ts.slots)

}
