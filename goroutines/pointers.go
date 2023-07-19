package main

import (
	"fmt"
)

func main() {
	a := 4
	f := &a
	fmt.Println(*f, f)
	fmt.Printf("%T\n", f)
	*f = 10
	fmt.Println(a, &a)
	squareAdd(&a)
	fmt.Println(a, &a)
}

func squareAdd(p *int) {
	g := 12
	p = &g
	fmt.Println(*p, p, &p)
}

type person struct {
	name string
	age  uint
}

func initPerson() *person {
	m := person{name: "john", age: 23}
	return &m
}
