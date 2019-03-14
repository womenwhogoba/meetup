package main

import (
	"fmt"
	"math" 
)

// START OMIT
type Vertice struct {
	X, Y float64
}

func (v Vertice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertice) Agrandar(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertice{3, 4}
	fmt.Println(v.Abs())
	v.Agrandar(10)
	fmt.Println(v.Abs())
}
// END OMIT
