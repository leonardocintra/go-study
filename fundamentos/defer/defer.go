package main

import (
	"fmt"
)

func obterValorAprovado(numero int) int {
	defer fmt.Println("FIM!")
	defer fmt.Println("Finalizou mesmo")

	if numero >= 500 {
		fmt.Println("Valor ALTO")
		return 500
	}

	fmt.Println("Valor BAIXO")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(600))
}
