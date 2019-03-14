package main

import "fmt"

func main() {

	nombres := [4]string{"John", "Paul", "George", "Ringo",}

	fmt.Println(nombres)

	a := nombres[0:2] // Creo un slice con los 2 primeros nombres
	b := nombres[1:3] // Creo un slice con el segundo y tercer nombre
	fmt.Println(a, b)

	b[0] = "XXX" // Modifico el primer nombre del slice b
	fmt.Println(a, b)
	fmt.Println(nombres)
}
