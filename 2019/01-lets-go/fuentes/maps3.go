package main

import "fmt"

type Vertice struct {
	Lat, Long float64
}
// START OMIT
var m = map[string]Vertice{
	"Bell Labs": Vertice{
		40.68433, -74.39967,
	},
	"Google": Vertice{
		37.42202, -122.08408,
	},
}

func main() {

	// agrego un nuevo Vertice
	m["Amazon"] = Vertice{-34.601156, -58.368232}

	elemento, existe := m["Amazonnn"] // pregunto por clave inexistente
	fmt.Printf("elemento: %v - existe: %v\n", elemento, existe)

	delete(m, "Bell Labssss") // elimino clave inexistente

	for clave, valor := range(m) {
		fmt.Println(clave, valor)	
	}
}
// END OMIT