package main

import "fmt"
// START OMIT
func main() {
	fmt.Println("contando")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("ok")
}
// END OMIT



