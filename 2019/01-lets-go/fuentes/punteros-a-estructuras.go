package main

import "fmt"

type Vertice struct {
	X int
	Y int
}

func main() {
	v := Vertice{1, 2}
	p := &v
	p.X = 10 // No es necesario hacer (*p).X // HL
	fmt.Println(v)
}

