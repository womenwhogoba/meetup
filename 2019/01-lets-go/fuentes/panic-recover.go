package main

import "fmt"
// START OMIT
func main() {
	f()
	fmt.Println("Salio bien de f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se recupera en f", r)
		}
	}()
	fmt.Println("Invocando a g.")
	g(0)
	fmt.Println("Salio bein de g.")
}
func g(i int) {
	if i > 3 {
		fmt.Println("PANICO!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer en g", i)
	fmt.Println("Hola, estoy en g", i)
	g(i + 1)
}
// END OMIT