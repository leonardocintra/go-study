package main

import (
	"fmt"
)

func main() {
	funcionarioPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva":     398.03,
			"Gustavo Rodrigues": 329.02,
		},
		"J": {
			"Joao Marcos": 200,
			"Juliana":     2020.32,
		},
		"P": {
			"Pedro Jose": 329,
			"Paulo":      984.3,
		},
	}

	delete(funcionarioPorLetra, "P")

	for letra, funcs := range funcionarioPorLetra {
		fmt.Println(letra, funcs)
	}
}
