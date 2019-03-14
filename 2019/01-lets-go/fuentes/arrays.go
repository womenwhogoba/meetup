package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hola"
	a[1] = "mundo"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primos)
}
