package main

import "fmt"
// START OMIT
type Motor struct {
	Cilindros int
}
type Auto struct {
	Motor
	Marca string
}
func (m Motor) Sonar() {
	fmt.Printf("Mis %v cilindros hacen poco ruido\n", m.Cilindros)
}
func (a Auto) Andar() {
	fmt.Printf("Con mis %v cilindros ando bien rapido\n", a.Cilindros)
}
func main() {
	fitito := Auto{Motor: Motor{Cilindros: 4}, Marca: "Fiat"}

	fitito.Sonar()
	fitito.Andar()
}
// END OMIT