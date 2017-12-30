package main

import (
	"fmt"

	"github.com/leonardocintra/area"
)

func main() {
	fmt.Println(area.Circulo(5.9))
	fmt.Println(area.Retangulo(5.0, 2.9))
	// fmt.Println(area._TrianguloEq(4, 6)) nao acessa pois é privado
}
