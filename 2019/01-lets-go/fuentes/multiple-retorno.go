package main

import "fmt"

func intercambiar(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := intercambiar("women who", "go")
	fmt.Println(a, b)
}

