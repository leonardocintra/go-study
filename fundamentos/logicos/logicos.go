package main

import (
	"fmt"
)

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTV50 := trabalho1 && trabalho2
	comprarTV32 := trabalho1 != trabalho2 // ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTV50, comprarTV32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, false)
	fmt.Printf("TV 50: %t, TV32 %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
