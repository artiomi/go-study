package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats[string, float64](floats))
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type List[T any] struct {
	prev *List[T]
	val  T
}

func runList() {
	integers := List[int]{val: 10}
	println(integers)
}

func (l *List[T]) String() string {
	return fmt.Sprintf("val: %v, prev: %v", l.val, l.prev)
}

func (l *List[T]) add(element T) {
	if l.prev != nil {
		l.prev.add(l.val)
	}
	l.prev = &List[T]{val: l.val, prev: l.prev}
	l.val = element
}
