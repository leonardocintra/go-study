package main

import (
	"fmt"
)

// funcoes variaticas
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Media %.2f", media(3, 2, 5.8, 2, 3, 1, 8, 8.8, 8))
}
