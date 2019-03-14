package main

import "fmt"

type Vertice struct {
	X int
	Y int
}

func main() {
	v := Vertice{1, 2}
	v.X = 4
	fmt.Println(v)
}

