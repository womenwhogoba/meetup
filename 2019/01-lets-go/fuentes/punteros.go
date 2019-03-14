package main

import "fmt"

func main() {
	i := 42

	p := &i         // puntero a i
	fmt.Println(*p) // Accedo a i a través del puntero
	*p = 21         // Modifico i a través del puntero
	fmt.Println(i)  // i fue modificado
}
