package main

import (
	"fmt"
)

func main() {
	// homogenea (mesmo tipo) e estatica (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7, 9.9, 4
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f \n", media)
}
