package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	// maps devem ser inicializados

	aprovados := make(map[int]string)
	aprovados[12345678] = "Maria"
	aprovados[11123232] = "Leonardo"
	aprovados[39899899] = "Ana"
	aprovados[33398392] = "ROnaldo"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[39899899])
	delete(aprovados, 39899899)
	fmt.Println(aprovados[39899899])
}
