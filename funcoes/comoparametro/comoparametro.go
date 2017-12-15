package main

import (
	"fmt"
)

func multiplicacao(a, b int) int {
	return a * b
}

func executar(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := executar(multiplicacao, 3, 9)
	fmt.Println(resultado)
}
