package main

import "fmt"

var empresas map[string]int

func main() {
	empresas = make(map[string]int)
	
	empresas["Grupo Esfera"] = 40
	empresas["Mercado Libre"] = 40000
	
	fmt.Println(empresas)
	fmt.Println(empresas["Grupo Esfera"])
}
