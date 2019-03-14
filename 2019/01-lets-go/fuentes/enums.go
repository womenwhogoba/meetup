package main

import "fmt"

type Operador int

const (
	Equals Operador = iota
	LessThan
	Like
)

func main() {
	fmt.Println("valor del operador Equals: ", Equals)
	fmt.Println("valor del operador Equals: ", LessThan)
	fmt.Println("valor del operador Equals: ", Like)
}
