package main

import "fmt"

func main() {
	var s = make([]int, 0)
	imprimirSlice(s)

	s = append(s, 0)
	imprimirSlice(s)

	s = append(s, 1)
	imprimirSlice(s)

	s = append(s, 2, 3, 4)
	imprimirSlice(s)
}

func imprimirSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
