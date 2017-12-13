package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6} // compilador conta a quantidade!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	// para nao ser obrigado a usar o indice "i" usar o _
	for _, num := range numeros {
		fmt.Println(num)
	}

	// se nao por o _ nem o i o for printa o INDICE!!
	for num := range numeros {
		fmt.Println(num)
	}

}
