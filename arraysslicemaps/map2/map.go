package main

import (
	"fmt"
)

func main() {
	funcionarioSalario := map[string]float64{
		"Leonardo":       123.35,
		"Gabriela":       3983.32,
		"Pedro Junior":   123.39,
		"Juliana Cintra": 3298.11,
	}

	funcionarioSalario["Rafale Silva"] = 1230.32
	fmt.Println(funcionarioSalario)
	delete(funcionarioSalario, "Indexistente")

	for nome, salario := range funcionarioSalario {
		fmt.Println(nome, salario)
	}
}
