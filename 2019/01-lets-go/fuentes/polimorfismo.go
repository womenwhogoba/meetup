package main

import "fmt"

type Perro struct {
}

type Gato struct {
}

// START OMIT

type Animal interface {
	EmitirSonido() string
}

func (p *Perro) EmitirSonido() string {
	return "guau"
}

func (g *Gato) EmitirSonido() string {
	return "miau"
}

func main() {
	var animales []Animal
	gato := &Gato{}
	perro := &Perro{}
	animales = append(animales, gato, perro)
	
	for _,animal := range(animales) {
		fmt.Println(animal.EmitirSonido())
	}
}
// END OMIT
