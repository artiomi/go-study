package main

import (
	"fmt"
)

func main() {
	//mapTest()
	slice()
}

func mapTest() {
	hits := make(map[string]map[string]int)
	add(hits, "/doc/", "au")
	fmt.Println(hits)

	var m = make(map[string]int, 12)
	m["a"] = 10
	m["b"] = 12
	fmt.Printf("map: %v, len:%v\n", m, len(m))
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
func slice() {

	b := [...]string{"John", "jane"}
	fmt.Println(b, cap(b))
	var iBuffer [10]int
	var slice = iBuffer[0:0]
	var slice2 = slice[0:3:6]
	for i := 0; i < 20; i++ {
		fmt.Printf("Slice cap %v and len %v\n", cap(slice), len(slice))
		if cap(slice) == len(slice) {
			fmt.Println("slice is full!")
			break
		}
		slice = Extend(slice, i)
		fmt.Println(slice)
		fmt.Println(iBuffer)
	}
	fmt.Println("------------------")
	fmt.Println(slice)
	fmt.Println(iBuffer)
	fmt.Println(slice2)

	slice2[2] = 99

	fmt.Println(slice)
	fmt.Println(iBuffer)
	fmt.Println(slice2)
}
func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func compute(fn func(float64, float64) float64, fn2 func()) float64 {
	fn2()
	return fn(3, 4)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {

	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
